local split = require("util.split")

describe("split", function()

  describe("get_last_value", function()
    it("splits value of an upstream variable and returns last value", function()
      for _, case in ipairs({{"127.0.0.1:26157 : 127.0.0.1:26158", "127.0.0.1:26158"},
                             {"127.0.0.1:26157, 127.0.0.1:26158", "127.0.0.1:26158"},
                             {"127.0.0.1:26158", "127.0.0.1:26158"}}) do
        local last = split.get_last_value(case[1])
        assert.equal(case[2], last)
      end
    end)
  end)

  describe("split_string", function()

    it("returns empty array if input string is empty", function()
      local splits, len = split.split_string("", ",")
      assert.equal(0, len)
      assert.is.truthy(splits)
    end)

    it("returns empty array if input string is nil", function()
      local splits, len = split.split_string(nil, ",")
      assert.equal(0, len)
      assert.is.truthy(splits)
    end)

    it("returns empty array if delimiter is empty", function()
      local splits, len = split.split_string("1,2", "")
      assert.equal(0, len)
      assert.is.truthy(splits)
    end)

    it("returns empty array delimiter is nil", function()
      local splits, len = split.split_string("1,2", nil)
      assert.equal(0, len)
      assert.is.truthy(splits)
    end)

    it("returns array of 1 value if input string is not a list", function()
      local splits, len = split.split_string("123", ",")
      assert.equal(1, len)
      assert.equal("123", splits[1])
    end)

    it("returns array of values extracted from the input string", function()
      local splits, len = split.split_string("1,2,3", ",")
      assert.equal(3, len)
      assert.equal("1", splits[1])
      assert.equal("2", splits[2])
      assert.equal("3", splits[3])
    end)

  end)
end)
