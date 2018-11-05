package msg

type MsgKeyExchange struct {
	SendSeed   int64
	ReciveSeed int64
}

type MsgHello struct {
	Data string
}
