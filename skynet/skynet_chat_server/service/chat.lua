local skynet = require "skynet"
require "skynet.manager"	-- import skynet.register
local mc = require "skynet.multicast"
local datacenter = require "skynet.datacenter"

local history_length = 10
local msg_db = {}
local msg_channel

local command = {}

function command.QUERY_MSG()
	return msg_db
end

function command.SEND_MSG(name, msg)
	local msg_item = {name=name, msg=msg}
	table.insert(msg_db, msg_item)
	print(string.format( "%s send msg：%s",name, msg))

	if #(msg_db) > history_length then
		table.remove(msg_db, 1)
	end
	-- 推送消息
	if msg_channel then
		print("publish to channel")
		msg_channel:publish(msg_item)
	end	
	return nil
end

skynet.start(function()
	skynet.dispatch("lua", function(session, address, cmd, ...)
		cmd = cmd:upper()
		if cmd == "PING" then
			assert(session == 0)
			local str = (...)
			if #str > 20 then
				str = str:sub(1,20) .. "...(" .. #str .. ")"
			end
			skynet.error(string.format("%s ping %s", skynet.address(address), str))
			return
		end
		local f = command[cmd]
		if f then
			skynet.ret(skynet.pack(f(...)))
		else
			error(string.format("Unknown command %s", tostring(cmd)))
		end
	end)

	-- 创建一个频道，成功创建后，channel.channel 是这个频道的 id
	local channel = mc.new()
	print("create msg channel id:", channel.channel)
	msg_channel = channel

	-- 加入到数据中心
	datacenter.set("MSG_CHANNEL", channel.channel)

--	skynet.traceproto("lua", false)	-- true off tracelog
	skynet.register "SIMPLECHAT"
end)
