local ngx           = ngx
local string        = string
local string_len    = string.len
local string_format = string.format
local pairs         = pairs
local ipairs        = ipairs
local tonumber      = tonumber
local getmetatable  = getmetatable
local type          = type
local next          = next
local table         = table
local re_gmatch     = ngx.re.gmatch

local _M = {}

function _M.get_nodes(endpoints)
  local nodes = {}
  local weight = 1

  for _, endpoint in pairs(endpoints) do
    local endpoint_string = endpoint.address .. ":" .. endpoint.port
    nodes[endpoint_string] = weight
  end

  return nodes
end

-- parse the compound variables, then call generate_var_value function
-- to parse into a string value.
function _M.parse_complex_value(complex_value)
    local reg = [[ (\\\$[0-9a-zA-Z_]+) | ]]     -- \$var
            .. [[ \$\{([0-9a-zA-Z_]+)\} | ]]    -- ${var}
            .. [[ \$([0-9a-zA-Z_]+) | ]]        -- $var
            .. [[ (\$|[^$\\]+) ]]               -- $ or text value
    local iterator, err = re_gmatch(complex_value, reg, "jiox")
    if not iterator then
        return nil, err
    end

    local v
    local t = {}
    while true do
        v, err = iterator()
        if err then
            return nil, err
        end

        if not v then
            break
        end

        table.insert(t, v)
    end

    return t
end

-- Parse the return value of function parse_complex_value
-- into a string value
function _M.generate_var_value(data)
  if data == nil then
    return ""
  end

  local t = {}
  for _, value in ipairs(data) do
    local var_name = value[2] or value[3]
    if var_name then
      if var_name:match("^%d+$") then
        var_name = tonumber(var_name)
      end
      table.insert(t, ngx.var[var_name])
    else
      table.insert(t, value[1] or value[4])
    end
  end

  return table.concat(t, "")
end

-- normalize_endpoints takes endpoints as an array of endpoint objects
-- and returns a table where keys are string that's
-- endpoint.address .. ":" .. endpoint.port and values are all true
local function normalize_endpoints(endpoints)
  local normalized_endpoints = {}

  for _, endpoint in pairs(endpoints) do
    local endpoint_string = string_format("%s:%s", endpoint.address, endpoint.port)
    normalized_endpoints[endpoint_string] = true
  end

  return normalized_endpoints
end

-- diff_endpoints compares old and new
-- and as a first argument returns what endpoints are in new
-- but are not in old, and as a second argument it returns
-- what endpoints are in old but are in new.
-- Both return values are normalized (ip:port).
function _M.diff_endpoints(old, new)
  local endpoints_added, endpoints_removed = {}, {}
  local normalized_old = normalize_endpoints(old)
  local normalized_new = normalize_endpoints(new)

  for endpoint_string, _ in pairs(normalized_old) do
    if not normalized_new[endpoint_string] then
      table.insert(endpoints_removed, endpoint_string)
    end
  end

  for endpoint_string, _ in pairs(normalized_new) do
    if not normalized_old[endpoint_string] then
      table.insert(endpoints_added, endpoint_string)
    end
  end

  return endpoints_added, endpoints_removed
end

-- this implementation is taken from
-- https://web.archive.org/web/20131225070434/http://snippets.
-- luacode.org/snippets/Deep_Comparison_of_Two_Values_3
-- and modified for use in this project
local function deep_compare(t1, t2, ignore_mt)
  local ty1 = type(t1)
  local ty2 = type(t2)
  if ty1 ~= ty2 then return false end
  -- non-table types can be directly compared
  if ty1 ~= 'table' and ty2 ~= 'table' then return t1 == t2 end
  -- as well as tables which have the metamethod __eq
  local mt = getmetatable(t1)
  if not ignore_mt and mt and mt.__eq then return t1 == t2 end
  for k1,v1 in pairs(t1) do
    local v2 = t2[k1]
    if v2 == nil or not deep_compare(v1,v2) then return false end
  end
  for k2,v2 in pairs(t2) do
    local v1 = t1[k2]
    if v1 == nil or not deep_compare(v1,v2) then return false end
  end
  return true
end
_M.deep_compare = deep_compare

function _M.is_blank(str)
  return str == nil or string_len(str) == 0
end

-- this implementation is taken from:
-- https://github.com/luafun/luafun/blob/master/fun.lua#L33
-- SHA: 04c99f9c393e54a604adde4b25b794f48104e0d0
local function deepcopy(orig)
  local orig_type = type(orig)
  local copy
  if orig_type == 'table' then
      copy = {}
      for orig_key, orig_value in next, orig, nil do
          copy[deepcopy(orig_key)] = deepcopy(orig_value)
      end
  else
      copy = orig
  end
  return copy
end
_M.deepcopy = deepcopy

local function tablelength(T)
  local count = 0
  for _ in pairs(T) do
      count = count + 1
  end
  return count
end
_M.tablelength = tablelength

-- replaces special character value a with value b for all occurrences in a
-- string
local function replace_special_char(str, a, b)
  return string.gsub(str, "%" .. a, b)
end
_M.replace_special_char = replace_special_char

local function request_addr()
    local req_headers = ngx.req.get_headers()
    local cli_addr = req_headers["x-forwarded-for"]
    if cli_addr == nil or string.len(cli_addr) == 0 or cli_addr == "unknown" then
        cli_addr = req_headers["Proxy-Client-IP"]
    end
    if cli_addr == nil or string.len(cli_addr) == 0 or cli_addr == "unknown" then
        cli_addr = req_headers["WL-Proxy-Client-IP"]
    end
    if cli_addr == nil or string.len(cli_addr) == 0 or cli_addr == "unknown" then
        cli_addr = ngx.var.remote_addr
    end
    -- Represents the situation through multiple agents, the first IP is the real IP of the client, and multiple IPs are divided according to','
    if cli_addr ~= nil and string.len(cli_addr) >15  then
        local pos  = string.find(cli_addr, ",", 1)
        cli_addr = string.sub(cli_addr,1,pos-1)
    end
    return cli_addr
end
_M.request_addr = request_addr

local function split(inputstr, sep)
    if sep == nil then
        sep ="%s"
    end
    local t={}
    for str in string.gmatch(inputstr,"([^"..sep.."]+)") do
        table.insert(t, str)
    end
    return t
end
_M.split = split

return _M
