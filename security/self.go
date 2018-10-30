package main

func Encrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	var p, i int
	var q byte
	for ks, bs := len(KEY), len(src); i < bs; i++ {
		p %= ks
		x := (src[i] ^ KEY[p]) + q
		p += 1
		q = x
		dst[i] = x
	}
	return dst
}

func Decrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	var p, i int
	var q byte
	for ks, bs := len(KEY), len(src); i < bs; i++ {
		p %= ks
		x := (src[i] - q) ^ KEY[p]
		p += 1
		q = src[i]
		dst[i] = x
	}
	return dst
}
