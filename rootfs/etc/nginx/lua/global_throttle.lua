local resty_global_throttle = require("resty.global_throttle")
local resty_ipmatcher = require("resty.ipmatcher")
local util = require("util")

local ngx = ngx
local ngx_exit = ngx.exit
local ngx_log = ngx.log
local ngx_ERR = ngx.ERR
local ngx_INFO = ngx.INFO

local ipairs = ipairs

local _M = {}

local DECISION_CACHE = ngx.shared.global_throttle_cache

-- it does not make sense to cache decision for too little time
-- the benefit of caching likely is negated if we cache for too little time
-- Lua Shared Dict's time resolution for expiry is 0.001.
local CACHE_THRESHOLD = 0.001

local DEFAULT_RAW_KEY = "remote_addr"

local function should_ignore_request(ignored_cidrs)
  if not ignored_cidrs or #ignored_cidrs == 0 then
    return false
  end

  local ignored_cidrs_matcher, err = resty_ipmatcher.new(ignored_cidrs)
  if not ignored_cidrs_matcher then
    ngx_log(ngx_ERR, "failed to initialize resty-ipmatcher: ", err)
    return false
  end

  local is_ignored
  is_ignored, err = ignored_cidrs_matcher:match(ngx.var.remote_addr)
  if err then
    ngx_log(ngx_ERR, "failed to match ip: '",
      ngx.var.remote_addr, "': ", err)
    return false
  end

  return is_ignored
end

local function has_value(tab, val)
  for _, value in ipairs(tab) do
    if value == val then
      return true
    end
  end
  return false
end

local function parse_header_based_objects(location_config)
  local request_headers = ngx.req.get_headers()
  for _,o in ipairs(location_config.header_based) do
    -- if we have value for header name, and the header is present in the request
    if ((o['header-name'] ~= "") and (request_headers[o['header-name']])) then
      -- check if the value matches the alternatives at "header-values"
      --  if match set limit and window_size and return
      --
      -- first match has precedence
      if has_value(o['header-values'], request_headers[o['header-name']]) then
        location_config['window_size'] = o['window']
        location_config['limit'] = o['limit']
        return
      end
    end
  end
end

local function is_enabled(config, location_config)
  if config.memcached.host == "" or config.memcached.port == 0 then
    return false
  end
  if #location_config.header_based > 0 then
    parse_header_based_objects(location_config)
  end
  if location_config.limit == 0 or
    location_config.window_size == 0 then
    return false
  end

  if should_ignore_request(location_config.ignored_cidrs) then
    return false
  end

  return true
end

local function get_namespaced_key_value(namespace, key_value)
  return namespace .. key_value
end

function _M.throttle(config, location_config)
  if not is_enabled(config, location_config) then
    return
  end

  local key_value = util.generate_var_value(location_config.key)
  if not key_value or key_value == "" then
    key_value = ngx.var[DEFAULT_RAW_KEY]
  end

  local namespaced_key_value =
    get_namespaced_key_value(location_config.namespace, key_value)

  local is_limit_exceeding = DECISION_CACHE:get(namespaced_key_value)
  if is_limit_exceeding then
    ngx.var.global_rate_limit_exceeding = "c"
    return ngx_exit(config.status_code)
  end

  local my_throttle, err = resty_global_throttle.new(
    location_config.namespace,
    location_config.limit,
    location_config.window_size,
    {
      provider = "memcached",
      host = config.memcached.host,
      port = config.memcached.port,
      connect_timeout = config.memcached.connect_timeout,
      max_idle_timeout = config.memcached.max_idle_timeout,
      pool_size = config.memcached.pool_size,
    }
  )
  if err then
    ngx.log(ngx.ERR, "faled to initialize resty_global_throttle: ", err)
    -- fail open
    return
  end

  local desired_delay, estimated_final_count
  estimated_final_count, desired_delay, err = my_throttle:process(key_value)
  if err then
    ngx.log(ngx.ERR, "error while processing key: ", err)
    -- fail open
    return
  end

  if desired_delay then
    if desired_delay > CACHE_THRESHOLD then
      local ok
      ok, err =
        DECISION_CACHE:safe_add(namespaced_key_value, true, desired_delay)
      if not ok then
        if err ~= "exists" then
          ngx_log(ngx_ERR, "failed to cache decision: ", err)
        end
      end
    end

    ngx.var.global_rate_limit_exceeding = "y"
    ngx_log(ngx_INFO, "limit is exceeding for ",
      location_config.namespace, "/", key_value,
      " with estimated_final_count: ", estimated_final_count)

    return ngx_exit(config.status_code)
  end
end

return _M
