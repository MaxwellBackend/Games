package main

func Encrypt(buf, encrypted []byte) {
	var p, i int
	var q byte
	for ks, bs := len(KEY), len(buf); i < bs; i++ {
		p %= ks
		x := (buf[i] ^ KEY[p]) + q
		p += 1
		q = x
		encrypted[i] = x
	}
}

func Decrypt(buf, decrypted []byte) {
	var p, i int
	var q byte
	for ks, bs := len(KEY), len(buf); i < bs; i++ {
		p %= ks
		x := (buf[i] - q) ^ KEY[p]
		p += 1
		q = buf[i]
		decrypted[i] = x
	}
}
