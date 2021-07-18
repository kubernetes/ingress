local ngx_balancer = require("ngx.balancer")
local cjson = require("cjson.safe")
local util = require("util")
local dns_lookup = require("util.dns").lookup
local configuration = require("configuration")
local round_robin = require("balancer.round_robin")
local chash = require("balancer.chash")
local chashsubset = require("balancer.chashsubset")
local sticky_balanced = require("balancer.sticky_balanced")
local sticky_persistent = require("balancer.sticky_persistent")
local ewma = require("balancer.ewma")
local string = string
local ipairs = ipairs
local table = table
local getmetatable = getmetatable
local tostring = tostring
local pairs = pairs
local math = math
local ngx = ngx

-- measured in seconds
-- for an Nginx worker to pick up the new list of upstream peers
-- it will take <the delay until controller POSTed the backend object to the
-- Nginx endpoint> + BACKENDS_SYNC_INTERVAL
local BACKENDS_SYNC_INTERVAL = 1
local BACKENDS_FORCE_SYNC_INTERVAL = 30

local DEFAULT_LB_ALG = "round_robin"
local IMPLEMENTATIONS = {
  round_robin = round_robin,
  chash = chash,
  chashsubset = chashsubset,
  sticky_balanced = sticky_balanced,
  sticky_persistent = sticky_persistent,
  ewma = ewma,
}

local PROHIBITED_LOCALHOST_PORT = configuration.prohibited_localhost_port or '10246'
local PROHIBITED_PEER_PATTERN = "^127.*:" .. PROHIBITED_LOCALHOST_PORT .. "$"

local _M = {}
local balancers = {}
local backends_with_external_name = {}
local backends_last_synced_at = 0

local function get_implementation(backend)
  local name = backend["load-balance"] or DEFAULT_LB_ALG

  if backend["sessionAffinityConfig"] and
     backend["sessionAffinityConfig"]["name"] == "cookie" then
    if backend["sessionAffinityConfig"]["mode"] == "persistent" then
      name = "sticky_persistent"
    else
      name = "sticky_balanced"
    end

  elseif backend["upstreamHashByConfig"] and
         backend["upstreamHashByConfig"]["upstream-hash-by"] then
    if backend["upstreamHashByConfig"]["upstream-hash-by-subset"] then
      name = "chashsubset"
    else
      name = "chash"
    end
  end

  local implementation = IMPLEMENTATIONS[name]
  if not implementation then
    ngx.log(ngx.WARN, backend["load-balance"], "is not supported, ",
            "falling back to ", DEFAULT_LB_ALG)
    implementation = IMPLEMENTATIONS[DEFAULT_LB_ALG]
  end

  return implementation
end

local function resolve_external_names(original_backend)
  local backend = util.deepcopy(original_backend)
  local endpoints = {}
  for _, endpoint in ipairs(backend.endpoints) do
    local ips = dns_lookup(endpoint.address)
    for _, ip in ipairs(ips) do
      table.insert(endpoints, { address = ip, port = endpoint.port })
    end
  end
  backend.endpoints = endpoints
  return backend
end

local function format_ipv6_endpoints(endpoints)
  local formatted_endpoints = {}
  for _, endpoint in ipairs(endpoints) do
    local formatted_endpoint = endpoint
    if not endpoint.address:match("^%d+.%d+.%d+.%d+$") then
      formatted_endpoint.address = string.format("[%s]", endpoint.address)
    end
    table.insert(formatted_endpoints, formatted_endpoint)
  end
  return formatted_endpoints
end

local function is_backend_with_external_name(backend)
  local serv_type = backend.service and backend.service.spec
                      and backend.service.spec["type"]
  return serv_type == "ExternalName"
end

local function sync_backend(backend)
  if not backend.endpoints or #backend.endpoints == 0 then
    balancers[backend.name] = nil
    return
  end

  if is_backend_with_external_name(backend) then
    backend = resolve_external_names(backend)
  end

  backend.endpoints = format_ipv6_endpoints(backend.endpoints)

  local implementation = get_implementation(backend)
  local balancer = balancers[backend.name]

  if not balancer then
    balancers[backend.name] = implementation:new(backend)
    return
  end

  -- every implementation is the metatable of its instances (see .new(...) functions)
  -- here we check if `balancer` is the instance of `implementation`
  -- if it is not then we deduce LB algorithm has changed for the backend
  if getmetatable(balancer) ~= implementation then
    ngx.log(ngx.INFO,
        string.format("LB algorithm changed from %s to %s, resetting the instance",
                      balancer.name, implementation.name))
    balancers[backend.name] = implementation:new(backend)
    return
  end

  balancer:sync(backend)
end

local function sync_backends_with_external_name()
  for _, backend_with_external_name in pairs(backends_with_external_name) do
    sync_backend(backend_with_external_name)
  end
end

local function sync_backends()
  local raw_backends_last_synced_at = configuration.get_raw_backends_last_synced_at()
  ngx.update_time()
  local current_timestamp = ngx.time()
  if current_timestamp - backends_last_synced_at < BACKENDS_FORCE_SYNC_INTERVAL
      and raw_backends_last_synced_at <= backends_last_synced_at then
    return
  end

  local backends_data = configuration.get_backends_data()
  if not backends_data then
    balancers = {}
    return
  end

  local new_backends, err = cjson.decode(backends_data)
  if not new_backends then
    ngx.log(ngx.ERR, "could not parse backends data: ", err)
    return
  end

  local balancers_to_keep = {}
  for _, new_backend in ipairs(new_backends) do
    if is_backend_with_external_name(new_backend) then
      local backend_with_external_name = util.deepcopy(new_backend)
      backends_with_external_name[backend_with_external_name.name] = backend_with_external_name
    else
      sync_backend(new_backend)
    end
    balancers_to_keep[new_backend.name] = true
  end

  for backend_name, _ in pairs(balancers) do
    if not balancers_to_keep[backend_name] then
      balancers[backend_name] = nil
      backends_with_external_name[backend_name] = nil
    end
  end
  backends_last_synced_at = raw_backends_last_synced_at
end

local function route_to_alternative_balancer(balancer)
  if balancer.is_affinitized(balancer) then
    -- If request is already affinitized to a primary balancer, keep the primary balancer.
    return false
  end

  if not balancer.alternative_backends then
    return false
  end

  -- TODO: support traffic shaping for n > 1 alternative backends
  local backend_name = balancer.alternative_backends[1]
  if not backend_name then
    ngx.log(ngx.ERR, "empty alternative backend")
    return false
  end

  local alternative_balancer = balancers[backend_name]
  if not alternative_balancer then
    ngx.log(ngx.ERR, "no alternative balancer for backend: ",
            tostring(backend_name))
    return false
  end

  if alternative_balancer.is_affinitized(alternative_balancer) then
    -- If request is affinitized to an alternative balancer, instruct caller to
    -- switch to alternative.
    return true
  end

  -- Use traffic shaping policy, if request didn't have affinity set.
  local traffic_shaping_policy =  alternative_balancer.traffic_shaping_policy
  if not traffic_shaping_policy then
    ngx.log(ngx.ERR, "traffic shaping policy is not set for balancer ",
            "of backend: ", tostring(backend_name))
    return false
  end

  local target_header = util.replace_special_char(traffic_shaping_policy.header,
                                                  "-", "_")
  local header = ngx.var["http_" .. target_header]
  if header then
    if traffic_shaping_policy.headerValue
	   and #traffic_shaping_policy.headerValue > 0 then
      if traffic_shaping_policy.headerValue == header then
        return true
      end
    elseif traffic_shaping_policy.headerPattern
       and #traffic_shaping_policy.headerPattern > 0 then
      local m, err = ngx.re.match(header, traffic_shaping_policy.headerPattern)
      if m then
        return true
      elseif  err then
          ngx.log(ngx.ERR, "error when matching canary-by-header-pattern: '",
                  traffic_shaping_policy.headerPattern, "', error: ", err)
          return false
      end
    elseif header == "always" then
      return true
    elseif header == "never" then
      return false
    end
  end

  local target_cookie = traffic_shaping_policy.cookie
  local cookie = ngx.var["cookie_" .. target_cookie]
  if cookie then
    if cookie == "always" then
      return true
    elseif cookie == "never" then
      return false
    end
  end

  if math.random(100) <= traffic_shaping_policy.weight then
    return true
  end

  return false
end

local function get_balancer_by_upstream_name(upstream_name)
  return balancers[upstream_name]
end

local function get_balancer()
  if ngx.ctx.balancer then
    return ngx.ctx.balancer
  end

  local backend_name = ngx.var.proxy_upstream_name

  local balancer = balancers[backend_name]
  if not balancer then
    return nil
  end

  if route_to_alternative_balancer(balancer) then
    local alternative_backend_name = balancer.alternative_backends[1]
    ngx.var.proxy_alternative_upstream_name = alternative_backend_name

    balancer = balancers[alternative_backend_name]
  end

  ngx.ctx.balancer = balancer

  return balancer
end

function _M.init_worker()
  -- when worker starts, sync non ExternalName backends without delay
  sync_backends()
  -- we call sync_backends_with_external_name in timer because for endpoints that require
  -- DNS resolution it needs to use socket which is not available in
  -- init_worker phase
  local ok, err = ngx.timer.at(0, sync_backends_with_external_name)
  if not ok then
    ngx.log(ngx.ERR, "failed to create timer: ", err)
  end

  ok, err = ngx.timer.every(BACKENDS_SYNC_INTERVAL, sync_backends)
  if not ok then
    ngx.log(ngx.ERR, "error when setting up timer.every for sync_backends: ", err)
  end
  ok, err = ngx.timer.every(BACKENDS_SYNC_INTERVAL, sync_backends_with_external_name)
  if not ok then
    ngx.log(ngx.ERR, "error when setting up timer.every for sync_backends_with_external_name: ",
            err)
  end
end

function _M.rewrite()
  local balancer = get_balancer()
  if not balancer then
    ngx.status = ngx.HTTP_SERVICE_UNAVAILABLE
    return ngx.exit(ngx.status)
  end
end

function _M.balance()
  local balancer = get_balancer()
  if not balancer then
    return
  end

  local peer = balancer:balance()
  if not peer then
    ngx.log(ngx.WARN, "no peer was returned, balancer: " .. balancer.name)
    return
  end

  if peer:match(PROHIBITED_PEER_PATTERN) then
    ngx.log(ngx.ERR, "attempted to proxy to self, balancer: ", balancer.name, ", peer: ", peer)
    return
  end

  ngx_balancer.set_more_tries(1)

  local ok, err = ngx_balancer.set_current_peer(peer)
  if not ok then
    ngx.log(ngx.ERR, "error while setting current upstream peer ", peer,
            ": ", err)
  end
end

function _M.log()
  local balancer = get_balancer()
  if not balancer then
    return
  end

  if not balancer.after_balance then
    return
  end

  balancer:after_balance()
end

setmetatable(_M, {__index = {
  get_implementation = get_implementation,
  sync_backend = sync_backend,
  route_to_alternative_balancer = route_to_alternative_balancer,
  get_balancer = get_balancer,
  get_balancer_by_upstream_name = get_balancer_by_upstream_name,
}})

return _M
