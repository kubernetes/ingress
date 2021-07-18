local ipairs = ipairs

local _M = {}

-- splits strings into host and port
local function parse_addr(addr)
  local _, _, host, port = addr:find("([^:]+):([^:]+)")
  if host and port then
    return {host=host, port=port}
  else
    return nil, "error in parsing upstream address!"
  end
end

function _M.get_first_value(var)
  local t = _M.split_upstream_var(var) or {}
  if #t == 0 then return nil end
  return t[1]
end

function _M.get_last_value(var)
  local t = _M.split_upstream_var(var) or {}
  if #t == 0 then return nil end
  return t[#t]
end

-- http://nginx.org/en/docs/http/ngx_http_upstream_module.html#example
-- CAVEAT: nginx is giving out : instead of , so the docs are wrong
-- 127.0.0.1:26157 : 127.0.0.1:26157 , ngx.var.upstream_addr
-- 200 : 200 , ngx.var.upstream_status
-- 0.00 : 0.00, ngx.var.upstream_response_time
function _M.split_upstream_var(var)
  if not var then
    return nil, nil
  end
  local t = {}
  for v in var:gmatch("[^%s|,]+") do
    if v ~= ":" then
      t[#t+1] = v
    end
  end
  return t
end

-- Splits an NGINX $upstream_addr and returns an array of tables
-- with a `host` and `port` key-value pair.
function _M.split_upstream_addr(addrs_str)
  if not addrs_str then
    return nil, nil
  end

  local addrs = _M.split_upstream_var(addrs_str)
  local host_and_ports = {}

  for _, v in ipairs(addrs) do
    local a, err = parse_addr(v)
    if err then
      return nil, err
    end
    host_and_ports[#host_and_ports+1] = a
  end
  if #host_and_ports == 0 then
    return nil, "no upstream addresses to parse!"
  end
  return host_and_ports
end

-- Splits string by delimiter. Returns array of parsed values and the length of the array.
function _M.split_string(what, delim)
  local result = {}
  local idx = 0

  if what and delim and delim ~= "" then
    for chunk in what:gmatch("([^" .. delim .. "]+)") do
      idx = idx + 1
      result[idx] = chunk
    end
  end

  return result, idx
end

return _M
