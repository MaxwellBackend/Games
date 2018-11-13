root = "/project/skynet/"

package.cpath = root.."luaclib/?.so"
package.path = root.."lualib/?.lua;"..root.."examples/?.lua"

if _VERSION ~= "Lua 5.3" then
	error "Use lua 5.3"
end

local socket = require "client.socket"
local proto = require "proto"
local sproto = require "sproto"

local host = sproto.new(proto.s2c):host "package"
local request = host:attach(sproto.new(proto.c2s))

local fd = assert(socket.connect("127.0.0.1", 8888))

local function send_package(fd, pack)
	local package = string.pack(">s2", pack)
	socket.send(fd, package)
end

local function unpack_package(text)
	local size = #text
	if size < 2 then
		return nil, text
	end
	local s = text:byte(1) * 256 + text:byte(2)
	if size < s+2 then
		return nil, text
	end

	return text:sub(3,2+s), text:sub(3+s)
end

local function recv_package(last)
	local result
	result, last = unpack_package(last)
	if result then
		return result, last
	end
	local r = socket.recv(fd)
	if not r then
		return nil, last
	end
	if r == "" then
		error "Server closed"
	end
	return unpack_package(last .. r)
end

local session = 0

local function send_request(name, args)
	session = session + 1
	local str = request(name, args, session)
	send_package(fd, str)
	print("Request:", session)
end

local last = ""

local function print_request(name, args)
	print("REQUEST", name)
	if args then
		for k,v in pairs(args) do
			print(k,v)
		end
	end
	if args and args.msg_list then -- 打印聊天历史信息
		for k,v in pairs(args.msg_list) do
			print(v.name, ":", v.msg)
		end
	end
end

local function print_response(session, args)
	print("RESPONSE", session)
	if args then
		for k,v in pairs(args) do
			print(k,v)
		end
		if args and args.msg_list then -- 打印聊天历史信息
			for k,v in pairs(args.msg_list) do
				print(v.name, ":", v.msg)
			end
		end
	end
end

local function print_package(t, ...)
	if t == "REQUEST" then
		print_request(...)
	else
		assert(t == "RESPONSE")
		print_response(...)
	end
end

local function dispatch_package()
	while true do
		local v
		v, last = recv_package(last)
		if not v then
			break
		end

		print_package(host:dispatch(v))
	end
end

function string.split(input, delimiter)
    input = tostring(input)
    delimiter = tostring(delimiter)
    if (delimiter=='') then return false end
    local pos,arr = 0, {}
    for st,sp in function() return string.find(input, delimiter, pos, true) end do
        table.insert(arr, string.sub(input, pos, st - 1))
        pos = sp + 1
    end
    table.insert(arr, string.sub(input, pos))
    return arr
end

send_request("handshake")
-- send_request("set", { what = "hello", value = "world" })
while true do
	dispatch_package()
	local line = socket.readstdin()
	if line then
		local arr = string.split(line, " ")
		if #(arr) > 0 then
			local cmd = arr[1]
			local msg = arr[2]
			if cmd and #(cmd)>0 then

				print(string.format( "cmd:%s, msg:%s", cmd, msg ))

				if cmd == "quit" then
					send_request("quit")
				else
					if cmd == "query_msg" then
						send_request("query_msg", {})
					elseif cmd =="send_msg" then
						local name = string.format( "user_%s", fd)
						print(name)
						send_request("send_msg", {name=name, msg=msg})
					end	
				end
			end
		end
	else
		socket.usleep(100)
	end
end