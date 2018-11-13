local skynet = require "skynet"
local socket = require "skynet.socket"
local sproto = require "sproto"
local sprotoloader = require "sprotoloader"
local mc = require "skynet.multicast"
local datacenter = require "skynet.datacenter"

local WATCHDOG
local host
local send_request

local CMD = {}
local REQUEST = {}
local client_fd

function REQUEST:query_msg()
	local r = skynet.call("SIMPLECHAT", "lua", "QUERY_MSG")
	return { msg_list = r }
end

function REQUEST:send_msg()
	local r = skynet.call("SIMPLECHAT", "lua", "SEND_MSG", self.name, self.msg)
	return { result = r }
end

function REQUEST:get()
	print("get", self.what)
	local r = skynet.call("SIMPLECHAT", "lua", "get", self.what)
	return { result = r }
end

function REQUEST:set()
	print("set", self.what, self.value)
	local r = skynet.call("SIMPLECHAT", "lua", "set", self.what, self.value)
end

function REQUEST:handshake()
	return { msg = "Welcome to skynet, I will send heartbeat every 5 sec." }
end

function REQUEST:quit()
	skynet.call(WATCHDOG, "lua", "close", client_fd)
end

local function request(name, args, response)
	local f = assert(REQUEST[name])
	local r = f(args)
	if response then
		return response(r)
	end
end

local function send_package(pack)
	local package = string.pack(">s2", pack)
	socket.write(client_fd, package)
end

skynet.register_protocol {
	name = "client",
	id = skynet.PTYPE_CLIENT,
	unpack = function (msg, sz)
		return host:dispatch(msg, sz)
	end,
	dispatch = function (fd, _, type, ...)
		assert(fd == client_fd)	-- You can use fd to reply message
		skynet.ignoreret()	-- session is fd, don't call skynet.ret
		skynet.trace()
		if type == "REQUEST" then
			local ok, result  = pcall(request, ...)
			if ok then
				if result then
					send_package(result)
				end
			else
				skynet.error(result)
			end
		else
			assert(type == "RESPONSE")
			error "This example doesn't support request client"
		end
	end
}

function recvChannel(channel, source, msg, ...)
	skynet.error("channel id:", channel.channel, "source:", skynet.address(source), "msg:", msg)
	local result = {msg_list={msg}}
	
	host = sprotoloader.load(1):host "package"
	send_request = host:attach(sprotoloader.load(2))

	send_package(send_request("push_msg", result))
end

function CMD.start(conf)
	local fd = conf.client
	local gate = conf.gate
	WATCHDOG = conf.watchdog
	-- slot 1,2 set at main.lua
	host = sprotoloader.load(1):host "package"
	send_request = host:attach(sprotoloader.load(2))
	skynet.fork(function()
		while true do
			send_package(send_request "heartbeat")
			skynet.sleep(5000)
		end
	end)

	client_fd = fd
	skynet.call(gate, "lua", "forward", fd)

	-- 订阅消息频道
	local msg_channel_id = datacenter.get("MSG_CHANNEL")
	print("agent get channel id:", msg_channel_id)
	channel = mc.new {
        channel = msg_channel_id, -- 绑定到消息频道
        dispatch = recvChannel,   -- 设置频道的消息处理函数
	}
	channel:subscribe()
end

function CMD.disconnect()
	-- todo: do something before exit
	skynet.exit()
end

skynet.start(function()
	skynet.dispatch("lua", function(_,_, command, ...)
		skynet.trace()
		local f = CMD[command]
		skynet.ret(skynet.pack(f(...)))
	end)
end)
