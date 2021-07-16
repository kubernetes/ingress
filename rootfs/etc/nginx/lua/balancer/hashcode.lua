local string = string
local _M = {}

function _M.str_hash_to_int(str)
    local h = 0
    local l = #str
    if l > 0 then
        local i = 0
        while i < l do
            h = 31 * h + string.byte(str, i + 1)
            i = i + 1
        end
    end
    return h
end

return _M
