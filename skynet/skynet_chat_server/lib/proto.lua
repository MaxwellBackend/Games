local sprotoparser = require "sprotoparser"

local proto = {}

proto.c2s = sprotoparser.parse [[
.package {
	type 0 : integer
	session 1 : integer
}

.MsgData {
	name 0 : string
	msg 1 : string
}

handshake 1 {
	response {
		msg 0  : string
	}
}

get 2 {
	request {
		what 0 : string
	}
	response {
		result 0 : string
	}
}

set 3 {
	request {
		what 0 : string
		value 1 : string
	}
}

quit 4 {}

query_msg 5 {
	request {
	}
	response {
		msg_list 0 : *MsgData
	}
}

send_msg 6 {
	request {
		name 0 : string
		msg 1 : string
	}
}

push_msg 7 {
	response {
		msg_list 0 : *MsgData
	}
}

]]

proto.s2c = sprotoparser.parse [[
.package {
	type 0 : integer
	session 1 : integer
}

.MsgData {
	name 0 : string
	msg 1 : string
}

heartbeat 1 {}

push_msg 7 {
	request {
		msg_list 0 : *MsgData
	}
	response {
	}
}

]]

return proto
