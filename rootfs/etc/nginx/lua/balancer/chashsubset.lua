-- Consistent hashing to a subset of nodes. Instead of returning the same node
-- always, we return the same subset always.

local resty_chash = require("resty.chash")
local util = require("util")
local ngx_log = ngx.log
local ngx_ERR = ngx.ERR
local setmetatable = setmetatable
local tostring = tostring
local math = math
local table = table
local pairs = pairs

local _M = { name = "chashsubset" }

local function build_subset_map(backend)
  local endpoints = {}
  local subset_map = {}
  local subsets = {}
  local subset_size = backend["upstreamHashByConfig"]["upstream-hash-by-subset-size"]

  for _, endpoint in pairs(backend.endpoints) do
    table.insert(endpoints, endpoint)
  end

  local set_count = math.ceil(#endpoints/subset_size)
  local node_count = set_count * subset_size

  -- if we don't have enough endpoints, we reuse endpoints in the last set to
  -- keep the same number on all of them.
  local j = 1
  for _ = #endpoints+1, node_count do
    table.insert(endpoints, endpoints[j])
    j = j+1
  end

  local k = 1
  for i = 1, set_count do
    local subset = {}
    local subset_id = "set" .. tostring(i)
    for _ = 1, subset_size do
      table.insert(subset, endpoints[k])
      k = k+1
    end
    subsets[subset_id] = subset
    subset_map[subset_id] = 1
  end

  return subset_map, subsets
end

function _M.new(self, backend)
  local subset_map, subsets = build_subset_map(backend)
  local complex_val, err =
    util.parse_complex_value(backend["upstreamHashByConfig"]["upstream-hash-by"])
  if err ~= nil then
    ngx_log(ngx_ERR, "could not parse the value of the upstream-hash-by: ", err)
  end

  local o = {
    instance = resty_chash:new(subset_map),
    hash_by = complex_val,
    subsets = subsets,
    current_endpoints = backend.endpoints
  }
  setmetatable(o, self)
  self.__index = self
  return o
end

function _M.is_affinitized()
  return false
end

function _M.balance(self)
  local key = util.generate_var_value(self.hash_by)
  local subset_id = self.instance:find(key)
  local endpoints = self.subsets[subset_id]
  local endpoint = endpoints[math.random(#endpoints)]
  return endpoint.address .. ":" .. endpoint.port
end

function _M.sync(self, backend)
  local subset_map

  local changed = not util.deep_compare(self.current_endpoints, backend.endpoints)
  if not changed then
    return
  end

  self.current_endpoints = backend.endpoints

  subset_map, self.subsets = build_subset_map(backend)

  self.instance:reinit(subset_map)

  return
end

return _M
