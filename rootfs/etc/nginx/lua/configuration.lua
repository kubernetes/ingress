local cjson = require("cjson.safe")
local util = require("util")

local io = io
local ngx = ngx
local tostring = tostring
local string = string
local table = table
local pairs = pairs

-- this is the Lua representation of Configuration struct in internal/ingress/types.go
local configuration_data = ngx.shared.configuration_data
local certificate_data = ngx.shared.certificate_data
local certificate_servers = ngx.shared.certificate_servers
local ocsp_response_cache = ngx.shared.ocsp_response_cache

local EMPTY_UID = "-1"
local BACKEND_BUCKET_SIZE = 10

local _M = {}

function _M.get_general_data()
  return configuration_data:get("general")
end

function _M.get_raw_backends_last_synced_at()
  local raw_backends_last_synced_at = configuration_data:get("raw_backends_last_synced_at")
  if raw_backends_last_synced_at == nil then
    raw_backends_last_synced_at = 1
  end
  return raw_backends_last_synced_at
end

local function fetch_request_body()
  ngx.req.read_body()
  local body = ngx.req.get_body_data()

  if not body then
    -- request body might've been written to tmp file if body > client_body_buffer_size
    local file_name = ngx.req.get_body_file()
    local file = io.open(file_name, "rb")

    if not file then
      return nil
    end

    body = file:read("*all")
    file:close()
  end

  return body
end

local function get_pem_cert(hostname)
  local uid = certificate_servers:get(hostname)
  if not uid then
    return nil
  end

  return certificate_data:get(uid)
end

local function handle_servers()
  if ngx.var.request_method ~= "POST" then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.print("Only POST requests are allowed!")
    return
  end

  local raw_configuration = fetch_request_body()

  local configuration, err = cjson.decode(raw_configuration)
  if not configuration then
    ngx.log(ngx.ERR, "could not parse configuration: ", err)
    ngx.status = ngx.HTTP_BAD_REQUEST
    return
  end

  local err_buf = {}

  for server, uid in pairs(configuration.servers) do
    if uid == EMPTY_UID then
      -- notice that we do not delete certificate corresponding to this server
      -- this is because a certificate can be used by multiple servers/hostnames
      certificate_servers:delete(server)
    else
      local success, set_err, forcible = certificate_servers:set(server, uid)
      if not success then
        local err_msg = string.format("error setting certificate for %s: %s\n",
          server, tostring(set_err))
        table.insert(err_buf, err_msg)
      end
      if forcible then
        local msg = string.format("certificate_servers dictionary is full, "
          .. "LRU entry has been removed to store %s", server)
        ngx.log(ngx.WARN, msg)
      end
    end
  end

  for uid, cert in pairs(configuration.certificates) do
    -- don't delete the cache here, certificate_data[uid] is not replaced yet.
    -- there is small chance that nginx worker still get the old certificate,
    -- then fetch and cache the old OCSP Response
    local old_cert = certificate_data:get(uid)
    local is_renew = (old_cert ~= nil and old_cert ~= cert)

    local success, set_err, forcible = certificate_data:set(uid, cert)
    if success then
        -- delete ocsp cache after certificate_data:set succeed
        if is_renew then
            ocsp_response_cache:delete(uid)
        end
    else
      local err_msg = string.format("error setting certificate for %s: %s\n",
        uid, tostring(set_err))
      table.insert(err_buf, err_msg)
    end
    if forcible then
      local msg = string.format("certificate_data dictionary is full, "
        .. "LRU entry has been removed to store %s", uid)
      ngx.log(ngx.WARN, msg)
    end
  end

  if #err_buf > 0 then
    ngx.log(ngx.ERR, table.concat(err_buf))
    ngx.status = ngx.HTTP_INTERNAL_SERVER_ERROR
    return
  end

  ngx.status = ngx.HTTP_CREATED
end

local function handle_general()
  if ngx.var.request_method == "GET" then
    ngx.status = ngx.HTTP_OK
    ngx.print(_M.get_general_data())
    return
  end

  if ngx.var.request_method ~= "POST" then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.print("Only POST and GET requests are allowed!")
    return
  end

  local config = fetch_request_body()

  local success, err = configuration_data:safe_set("general", config)
  if not success then
    ngx.status = ngx.HTTP_INTERNAL_SERVER_ERROR
    ngx.log(ngx.ERR, "error setting general config: " .. tostring(err))
    return
  end

  ngx.status = ngx.HTTP_CREATED
end

local function handle_certs()
  if ngx.var.request_method ~= "GET" then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.print("Only GET requests are allowed!")
    return
  end

  local query = ngx.req.get_uri_args()
  if not query["hostname"] then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.print("Hostname must be specified.")
    return
  end

  local key = get_pem_cert(query["hostname"])
  if key then
    ngx.status = ngx.HTTP_OK
    ngx.print(key)
    return
  else
    ngx.status = ngx.HTTP_NOT_FOUND
    ngx.print("No key associated with this hostname.")
    return
  end
end

local function get_backends_with_name(name)
  local backends_data = configuration_data:get(name)

  local backends, err = cjson.decode(backends_data)
  if not backends_data then
    ngx.log(ngx.ERR, "could not parse backends data: ", err)
    return
  end
  return backends
end

local function get_backend_bucket_names()
  local names = configuration_data:get("backend_bucket_names")
  if not names then
    return
  end
  local backend_bucket_names, err = cjson.decode(names)
  if not backend_bucket_names then
    ngx.log(ngx.ERR, "could not parse backend bucket names data: ", err)
    return
  end

  return backend_bucket_names
end

local function get_backends(is_init_call)
  local bucket_names = get_backend_bucket_names()
  if not bucket_names then
    ngx.log(ngx.WARN, "bucket name not found")
    return
  end
  local all_backends = {}
  for i, name in ipairs(bucket_names) do
    local backends = get_backends_with_name(name)
    for j, v in ipairs(backends) do
      table.insert(all_backends, v)
    end

    if not is_init_call then
      ngx.sleep(0)
    end
  end

  return all_backends
end

local function save_backends(backends)
  local backend_buckets = {}
  local backend_bucket_names = {}
  local backend_bucket_size = BACKEND_BUCKET_SIZE
  local tmp_backends = {}
  for i, v in ipairs(backends) do
    if table.getn(tmp_backends) >= BACKEND_BUCKET_SIZE then
       local bucket_key = string.format("bucket_%d", backend_bucket_size)
       local backends = util.deepcopy(tmp_backends)
       table.insert(backend_bucket_names, bucket_key)
       table.insert(backend_buckets, backends)
       tmp_backends = {}
       backend_bucket_size = backend_bucket_size + BACKEND_BUCKET_SIZE
    end
    table.insert(tmp_backends, v)
  end
  if table.getn(tmp_backends) > 0 then
    local bucket_key = string.format("bucket_%d", backend_bucket_size)
    local backends = util.deepcopy(tmp_backends)
    table.insert(backend_bucket_names, bucket_key)
    table.insert(backend_buckets, backends)
    tmp_backends = {}
  end

  for i, backends in ipairs(backend_buckets) do
    local new_backends, err = cjson.encode(backends)
    if not new_backends then
      ngx.log(ngx.ERR, "could not parse backends data: ", err)
      ngx.status = ngx.HTTP_BAD_REQUEST
      return
    end

    local backend_bucket_name = backend_bucket_names[i]
    local success, err = configuration_data:set(backend_bucket_name, new_backends)
    if not success then
      ngx.log(ngx.ERR, "dynamic-configuration: error updating configuration: " .. tostring(err))
      ngx.status = ngx.HTTP_BAD_REQUEST
      return
    end
  end

  local new_backend_names, err = cjson.encode(backend_bucket_names)
  if not new_backend_names then
    ngx.log(ngx.ERR, "could not parse backends_names data: ", err)
    ngx.status = ngx.HTTP_BAD_REQUEST
    return
  end

  local success, err = configuration_data:set("backend_bucket_names", new_backend_names)
  if not success then
    ngx.log(ngx.ERR, "dynamic-configuration: error updating configuration: " .. tostring(err))
    ngx.status = ngx.HTTP_BAD_REQUEST
    return
  end
end

local function handle_backends()
  if ngx.var.request_method == "GET" then
    local backneds = get_backends()
    local backneds_data, err = cjson.decode(backends)
    if not backneds_data then
      ngx.log(ngx.ERR, "could not parse backends data: ", err)
      ngx.status = ngx.HTTP_BAD_REQUEST
      return
    end
    ngx.status = ngx.HTTP_OK
    ngx.print(backneds_data)
    return
  end

  local backends = fetch_request_body()
  if not backends then
    ngx.log(ngx.ERR, "dynamic-configuration: unable to read valid request body")
    ngx.status = ngx.HTTP_BAD_REQUEST
    return
  end

  local new_backends, err = cjson.decode(backends)
  if not new_backends then
    ngx.log(ngx.ERR, "could not parse backends data: ", err)
    ngx.status = ngx.HTTP_BAD_REQUEST
    return
  end
  save_backends(new_backends)

  ngx.update_time()
  local raw_backends_last_synced_at = ngx.time()
  local success, err = configuration_data:set("raw_backends_last_synced_at", raw_backends_last_synced_at)
  if not success then
    ngx.log(ngx.ERR, "dynamic-configuration: error updating when backends sync, " ..
                     "new upstream peers waiting for force syncing: " .. tostring(err))
    ngx.status = ngx.HTTP_BAD_REQUEST
    return
  end

  ngx.status = ngx.HTTP_CREATED
end

function _M.call()
  if ngx.var.request_method ~= "POST" and ngx.var.request_method ~= "GET" then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.print("Only POST and GET requests are allowed!")
    return
  end

  if ngx.var.request_uri == "/configuration/servers" then
    handle_servers()
    return
  end

  if ngx.var.request_uri == "/configuration/general" then
    handle_general()
    return
  end

  if ngx.var.uri == "/configuration/certs" then
    handle_certs()
    return
  end

  if ngx.var.request_uri == "/configuration/backends" then
    handle_backends()
    return
  end

  ngx.status = ngx.HTTP_NOT_FOUND
  ngx.print("Not found!")
end

function _M.get_backends_data(is_init_call)
  return get_backends(is_init_call)
end

setmetatable(_M, {__index = { handle_servers = handle_servers }})

return _M
