package security

import (
	"math/big"
	"testing"
)

var bts = []byte("maxwell backend is a newbest team,do u know ? yes, i know,i know")
var tmp []byte
var dst []byte

var aesCipher ICipher
var desCipher ICipher
var rc4Cipher ICipher
var rsaCipher ICipher

func TestMain(m *testing.M) {
	x1, _ := DHExchange()
	key1 := DHKey(x1, big.NewInt(int64(123)))
	aesCipher = NewAesCipher(Itob(key1))
	desCipher = NewDesCipher(Itob(key1))
	rc4Cipher = NewRc4Cipher(Itob(key1))
	rsaCipher = NewRsaCipher("pem/server_public.pem", "pem/server_private.pem")
	m.Run()
}
