local util = require("util")

local function assert_request_rejected(config, location_config, opts)
  stub(ngx, "exit")

  local global_throttle = require_without_cache("global_throttle")
  assert.has_no.errors(function()
    global_throttle.throttle(config, location_config)
  end)

  assert.stub(ngx.exit).was_called_with(config.status_code)
  if opts.with_cache then
    assert.are.same("c", ngx.var.global_rate_limit_exceeding)
  else
    assert.are.same("y", ngx.var.global_rate_limit_exceeding)
  end
end

local function assert_request_not_rejected(config, location_config)
  stub(ngx, "exit")
  local cache_safe_add_spy = spy.on(ngx.shared.global_throttle_cache, "safe_add")

  local global_throttle = require_without_cache("global_throttle")
  assert.has_no.errors(function()
    global_throttle.throttle(config, location_config)
  end)

  assert.stub(ngx.exit).was_not_called()
  assert.is_nil(ngx.var.global_rate_limit_exceeding)
  assert.spy(cache_safe_add_spy).was_not_called()
end

local function assert_short_circuits(f)
  local cache_get_spy = spy.on(ngx.shared.global_throttle_cache, "get")

  local resty_global_throttle = require_without_cache("resty.global_throttle")
  local resty_global_throttle_new_spy = spy.on(resty_global_throttle, "new")

  local global_throttle = require_without_cache("global_throttle")

  f(global_throttle)

  assert.spy(resty_global_throttle_new_spy).was_not_called()
  assert.spy(cache_get_spy).was_not_called()
end

local function assert_fails_open(config, location_config, ...)
  stub(ngx, "exit")
  stub(ngx, "log")

  local global_throttle = require_without_cache("global_throttle")

  assert.has_no.errors(function()
    global_throttle.throttle(config, location_config)
  end)

  assert.stub(ngx.exit).was_not_called()
  assert.stub(ngx.log).was_called_with(ngx.ERR, ...)
  assert.is_nil(ngx.var.global_rate_limit_exceeding)
end

local function stub_resty_global_throttle_process(ret1, ret2, ret3, f)
  local resty_global_throttle = require_without_cache("resty.global_throttle")
  local resty_global_throttle_mock = {
    process = function(self, key) return ret1, ret2, ret3 end
  }
  stub(resty_global_throttle, "new", resty_global_throttle_mock)

  f()

  assert.stub(resty_global_throttle.new).was_called()
end

local function cache_rejection_decision(namespace, key_value, desired_delay)
  local namespaced_key_value = namespace .. key_value
  local ok, err = ngx.shared.global_throttle_cache:safe_add(namespaced_key_value, true, desired_delay)
  assert.is_nil(err)
  assert.is_true(ok)
  assert.is_true(ngx.shared.global_throttle_cache:get(namespaced_key_value))
end

describe("global_throttle", function()
  local snapshot

  local NAMESPACE = "31285d47b1504dcfbd6f12c46d769f6e"
  local LOCATION_CONFIG = {
    namespace = NAMESPACE,
    limit = 10,
    window_size = 60,
    key = {},
    ignored_cidrs = {},
    header_based = {},
  }
  local CONFIG = {
    memcached = {
      host = "memc.default.svc.cluster.local", port = 11211,
      connect_timeout = 50, max_idle_timeout = 10000, pool_size = 50,
    },
    status_code = 429,
  }

  before_each(function()
    snapshot = assert:snapshot()

    ngx.var = { remote_addr = "127.0.0.1", global_rate_limit_exceeding = nil, http_x_client_role = "root" }
    ngx.req = {
      get_headers = function() return {
        ["host"]          = "127.0.0.1",
        ["user-agent"]    = "lua-resty-http",
        ["x-client-role"] = "root"
      } end,
    }
  end)

  after_each(function()
    snapshot:revert()

    ngx.shared.global_throttle_cache:flush_all()
    reset_ngx()
  end)

  it("short circuits when memcached is not configured", function()
    assert_short_circuits(function(global_throttle)
      assert.has_no.errors(function()
        global_throttle.throttle({ memcached = { host = "", port = 0 } }, LOCATION_CONFIG)
      end)
    end)
  end)

  it("short circuits when limit or window_size is not configured", function()
    assert_short_circuits(function(global_throttle)
      local location_config_copy = util.deepcopy(LOCATION_CONFIG)
      location_config_copy.limit = 0
      assert.has_no.errors(function()
        global_throttle.throttle(CONFIG, location_config_copy)
      end)
    end)

    assert_short_circuits(function(global_throttle)
      local location_config_copy = util.deepcopy(LOCATION_CONFIG)
      location_config_copy.window_size = 0
      assert.has_no.errors(function()
        global_throttle.throttle(CONFIG, location_config_copy)
      end)
    end)
  end)

  it("short circuits when remote_addr is in ignored_cidrs", function()
    local global_throttle = require_without_cache("global_throttle")
    local location_config = util.deepcopy(LOCATION_CONFIG)
    location_config.ignored_cidrs = { ngx.var.remote_addr }
    assert_short_circuits(function(global_throttle)
      assert.has_no.errors(function()
        global_throttle.throttle(CONFIG, location_config)
      end)
    end)
  end)

  it("short circuits when X-Client-Role is 'admin' or 'root' in header_based", function()
    local global_throttle = require_without_cache("global_throttle")
    local location_config = util.deepcopy(LOCATION_CONFIG)
    location_config.header_based = { {["header-name"]="x-client-role", ["header-values"]={"admin", "root"}, ["limit"]=0, ["window_size"]=0} }
    assert_short_circuits(function(global_throttle)
      assert.has_no.errors(function()
        global_throttle.throttle(CONFIG, location_config)
      end)
    end)
  end)

  it("rejects when exceeding limit has already been cached", function()
    local key_value = "foo"
    local location_config = util.deepcopy(LOCATION_CONFIG)
    location_config.key = { { nil, nil, nil, key_value } }
    cache_rejection_decision(NAMESPACE, key_value, 0.5)

    assert_request_rejected(CONFIG, location_config, { with_cache = true })
  end)

  describe("when resty_global_throttle fails", function()
    it("fails open in case of initialization error", function()
      local too_long_namespace = ""
      for i=1,36,1 do
        too_long_namespace = too_long_namespace .. "a"
      end

      local location_config = util.deepcopy(LOCATION_CONFIG)
      location_config.namespace = too_long_namespace

      assert_fails_open(CONFIG, location_config, "faled to initialize resty_global_throttle: ", "'namespace' can be at most 35 characters")
    end)

    it("fails open in case of key processing error", function()
      stub_resty_global_throttle_process(nil, nil, "failed to process", function()
        assert_fails_open(CONFIG, LOCATION_CONFIG, "error while processing key: ", "failed to process")
      end)
    end)
  end)

  it("initializes resty_global_throttle with the right parameters", function()
    local resty_global_throttle = require_without_cache("resty.global_throttle")
    local resty_global_throttle_original_new = resty_global_throttle.new
    resty_global_throttle.new = function(namespace, limit, window_size, store_opts)
      local o, err = resty_global_throttle_original_new(namespace, limit, window_size, store_opts)
      if not o then
        return nil, err
      end
      o.process = function(self, key) return 1, nil, nil end

      local expected = LOCATION_CONFIG
      assert.are.same(expected.namespace, namespace)
      assert.are.same(expected.limit, limit)
      assert.are.same(expected.window_size, window_size)

      assert.are.same("memcached", store_opts.provider)
      assert.are.same(CONFIG.memcached.host, store_opts.host)
      assert.are.same(CONFIG.memcached.port, store_opts.port)
      assert.are.same(CONFIG.memcached.connect_timeout, store_opts.connect_timeout)
      assert.are.same(CONFIG.memcached.max_idle_timeout, store_opts.max_idle_timeout)
      assert.are.same(CONFIG.memcached.pool_size, store_opts.pool_size)

      return o, nil
    end
    local resty_global_throttle_new_spy = spy.on(resty_global_throttle, "new")

    local global_throttle = require_without_cache("global_throttle")

    assert.has_no.errors(function()
      global_throttle.throttle(CONFIG, LOCATION_CONFIG)
    end)

    assert.spy(resty_global_throttle_new_spy).was_called()
  end)

  it("rejects request and caches decision when limit is exceeding after processing a key", function()
    local desired_delay = 0.015

    stub_resty_global_throttle_process(LOCATION_CONFIG.limit + 1, desired_delay, nil, function()
      assert_request_rejected(CONFIG, LOCATION_CONFIG, { with_cache = false })

      local cache_key = LOCATION_CONFIG.namespace .. ngx.var.remote_addr
      assert.is_true(ngx.shared.global_throttle_cache:get(cache_key))

      -- we assume it won't take more than this  after caching
      -- until we execute the assertion below
      local delta = 0.001
      local ttl = ngx.shared.global_throttle_cache:ttl(cache_key)
      assert.is_true(ttl > desired_delay - delta)
      assert.is_true(ttl <= desired_delay)
    end)
  end)

  it("rejects request and skip caching of decision when limit is exceeding after processing a key but desired delay is lower than the threshold", function()
    local desired_delay = 0.0009

    stub_resty_global_throttle_process(LOCATION_CONFIG.limit, desired_delay, nil, function()
      assert_request_rejected(CONFIG, LOCATION_CONFIG, { with_cache = false })

      local cache_key = LOCATION_CONFIG.namespace .. ngx.var.remote_addr
      assert.is_nil(ngx.shared.global_throttle_cache:get(cache_key))
    end)
  end)

  it("allows the request when limit is not exceeding after processing a key", function()
    stub_resty_global_throttle_process(LOCATION_CONFIG.limit - 3, nil, nil,
      function()
        assert_request_not_rejected(CONFIG, LOCATION_CONFIG)
      end
    )
  end)

  it("rejects with custom status code", function()
    cache_rejection_decision(NAMESPACE, ngx.var.remote_addr, 0.3)
    local config = util.deepcopy(CONFIG)
    config.status_code = 503
    assert_request_rejected(config, LOCATION_CONFIG, { with_cache = true })
  end)
end)
