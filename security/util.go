package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func RandUint(min uint, max uint) uint {
	if min > max {
		min, max = max, min
	}

	return (uint(rand.Int63()) % (max - min + 1)) + min
}

func PubKey() (*rsa.PublicKey, error) {
	data := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv716bDn5TJ8et5pf9cB+
ibzXcAG+6qObSm9ExX6GFQvp70sSGa0MnQyjwMMCBZ0UC5bfVnwvixG89M30SEvq
w9hUHVjfRa8llFaU2cq9oEVVbVAHWLWHT6wpdvFYn8NIjwDEjoFPLi7Pc2q2qjbK
CP+bxzL6k45qidu/3ii27C0YcdCBa2kil4jDunIth1hXYHJHi4I9krE6ceM12iN+
vUX8jw5uJqBvNyEZHmwbhjVFfdrNh8nwIHM5HyXB4m8KnG6F++oxD5PJHtP1J6FT
G71wxE16P5t2P0aFYoq8fU5stDtcnXBizaqP7Vex7O9iH0Y8BKUfbzWo1dSQ/yvm
KQIDAQAB
-----END PUBLIC KEY-----`
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		return nil, errors.New("Bad public key")
	}
	if got, want := block.Type, "PUBLIC KEY"; got != want {
		return nil, errors.New("Unknown key type")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PublicKey), nil
}

func PriKey() (*rsa.PrivateKey, error) {
	data := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAlr24ZlJC1oDV8h28SR7bIC3aNPB13e6p+Qly8WbipC54uTBo
Esp1ZrZjjJusyO3feSz5eclHx2EOqibBf8wshAbCFsDdSnlZ+S4sp6DU5soaiBVv
Qc750DODbIun3pPeu2QuPr9XG5PsK5awbYB5uY+idn86TBhG/fnsgIQnUFNRDEWJ
yor9BFzv/GFi4rT1yKBUCKmbfbDP5yF2aZ05pgdIamPq5/O1OR8y757WuuHA8y8/
aGMzdAHZFQpkmQMYW92kS4/4D/3+gPynKrphmHmiV9DdltsPDr2Sb78QyND+VhU+
L7z1fpRT1+IMa4xt8out7ZE8DFnsPsvahVzNQQIDAQABAoIBAA7NFCUf4KY1Qkyi
tNEtMXDob6uRDNG81H26esnPQ2PBmGL+qMlnjb2HHSE6S+3yxWWdj+VwaRbfpdP5
ODi3MIvKKhaX22KrpbT83q+a8Cy1HYPNLv50Nz1FQZtP3yzQsicOQCgzG3d+v8rL
4gjue4VH4DjoZy3/mEDp4FoCH+YYEaCxF1HRQhsQhXH1nuGa3Io5TvLZBvq7WRd5
oEIOVrm+9UfjE0HAwieVpJ+sg//SLUMzrqdWUlWDfeeaaQ9jqImKlmIzS8R/J1j0
9Nd+teeHtv+dyWphg0bsG92TUHlCx9389shjfl80m3JaVujkHvUu1zFDkncXU/DF
+I11xA0CgYEAyBbZ2xPQUDs4Y9wW1jxlAbp7v8cSW/zpN/AWYnfgE0jU0f/Zjofp
dxwnjsNifGQ7dVqPdNRdkiKqdCi0lDHaLXLO/c6YGefed8YfqcdhEAe9qjig9715
dTNqVk1gWCYaBGtM/akSMcBW0px8NjACc9lD9j/OMp+eoYcQasPr/P8CgYEAwNzO
djnyz59OdTE8ZKjKAAa/19ciS2hXGIyEYptY8lNDm5RCEgdTFoQXa0+i9v0DQ3cr
WpcBtGt8QYNWOazer772YtOgi+9G9PZx+6jTrGs8DDLjULzmXl1by+upeHGSxOV1
RmZaS1VqEzIEhL+8G9yo5ByQuQmuewyhpQHV9b8CgYEArMOmSdY62Pu5PegMQ+ET
6cEmFroBrTGcAbOo8E7HSH9rWwOJytqpiDHkKg2kXPmCqVqPYXX4cTTDbtayzP6a
fwGRnm6iwEOHwG1ua37+3QOCDZclzaNpWvwRBgF7fcEwV756VM+GKTUFAochbPxZ
McRYPShslyNhqeDgGP+gJOsCgYBS24poCKVkxdBzUN/dbxa5xCyPkoHNW9pvVY9a
vdsM4PDe2mc4T6VGuyFEnUYqTrEkPwKFPPMijDsp+j8FMFvLrYM1Gi6TcA1QGMEH
mZS3Mqrtor8+ZKmkSP+zMF/yuviPVvV3dWm9i8iyzn5aefVbEwYMWqlileI8Q0T/
GzbNCQKBgAVva+RY9VmSMYQggUrrDktywMAoMG4gE8gcysHk2W1ykfV3SyV0u6x8
o3yDMCAlTh49lDPcRrBkzYVSlHC2y4qgvnTF77nmaIUk/jIrv/G2oZ8Co7VjPsvT
wJqoez5uATg+8CCdoWCG0qsj+mlqpeeu2kCfPCyySZQs1wNofcju
-----END RSA PRIVATE KEY-----`
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		return nil, errors.New("Bad private key")
	}
	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		return nil, errors.New("Unknown key type")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// Diffie-Hellman Key-exchange
var (
	rng         = rand.New(rand.NewSource(time.Now().UnixNano()))
	DH1BASE     = big.NewInt(3)
	DH1PRIME, _ = big.NewInt(0).SetString("0x7FFFFFC3", 0)
	MAXINT64    = big.NewInt(math.MaxInt64)
)

func DHExchange() (*big.Int, *big.Int) {
	secret := big.NewInt(0).Rand(rng, MAXINT64)
	modpower := big.NewInt(0).Exp(DH1BASE, secret, DH1PRIME)
	return secret, modpower
}

func DHKey(SECRET, MODPOWER *big.Int) *big.Int {
	key := big.NewInt(0).Exp(MODPOWER, SECRET, DH1PRIME)
	return key
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
