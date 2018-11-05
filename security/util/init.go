package util

import (
	"log"
	"math/big"
)

var bts = []byte("maxwell backend is a newbest team,do u know ? yes, i know,i know")
var tmp []byte
var dst []byte

var aesCipher ICipher
var desCipher ICipher
var rc4Cipher ICipher

func init() {
	x1, _ := DHExchange()
	key1 := DHKey(x1, big.NewInt(int64(123)))

	aesCipher = NewAesCipher(Itob(key1))
	desCipher = NewDesCipher(Itob(key1))
	rc4Cipher = NewRc4Cipher(Itob(key1))
	log.Printf("init complete")
}
