local util = require("util")

local string_format = string.format
local ngx_log = ngx.log
local INFO = ngx.INFO
local setmetatable = setmetatable

local _M = {}

function _M.new(self, o)
  o = o or {}
  setmetatable(o, self)
  self.__index = self
  return o
end

function _M.is_affinitized()
  return false
end

function _M.sync(self, backend)
  self.traffic_shaping_policy = backend.trafficShapingPolicy
  self.alternative_backends = backend.alternativeBackends

  local nodes = util.get_nodes(backend.endpoints)
  local changed = not util.deep_compare(self.instance.nodes, nodes)
  if not changed then
    return
  end

  ngx_log(INFO, string_format("[%s] nodes have changed for backend %s", self.name, backend.name))

  self.instance:reinit(nodes)
end

return _M
