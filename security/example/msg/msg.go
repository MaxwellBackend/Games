package msg

type MsgHandShake struct {
	Key []byte
}

type MsgKeyExchange struct {
	SendSeed   int64
	ReciveSeed int64
}

type MsgHello struct {
	Data string
}
