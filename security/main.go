package main

import (
	"fmt"
	"time"
)

func main() {
	var bts []byte
	var tmp []byte
	var dst []byte
	time := func(exe func(), log string) {
		start := time.Now()
		exe()
		end := time.Now()
		fmt.Println(log, len(bts), len(dst), end.Sub(start))
	}
	fill := func(text string, length int) {
		if length == 0 {
			length = len(text)
		}
		bts = make([]byte, length)
		dst = make([]byte, length)
		tmp = make([]byte, length)
		if text != "" {
			copy(bts, []byte(text))
		} else {
			for i := range bts {
				bts[i] = byte(RandUint(0, 255))
			}
		}

		time(func() {
			Encrypt(bts[:], tmp[:])
		}, "self encrypt")

		time(func() {
			Decrypt(tmp[:], dst[:])
		}, "self decrypt")

		time(func() {
			AesEncrypt(bts[:], tmp[:])
		}, "aes encrypt")

		time(func() {
			AesDecrypt(tmp[:], dst[:])
		}, "aes decrypt")

		time(func() {
			DesEncrypt(bts[:], tmp[:])
		}, "des encrypt")

		time(func() {
			DesDecrypt(tmp[:], dst[:])
		}, "des decrypt")

		time(func() {
			Rc4Encrypt(bts[:], tmp[:])
		}, "rc4 encrypt")

		time(func() {
			Rc4Decrypt(tmp[:], dst[:])
		}, "rc4 decrypt")

		fmt.Println()
	}

	fill("this project is very good!", 0)
	fill("hello,this project is very good!", 0)
	fill("", 24)
	fill("", 255)
	fill("", 1025)
	fill("", 2045)
	fill("", 4045)
}
