## 交换key
	OAEP  最佳非对称加密填充（OAEP）是一个通常和RSA一起使用的填充方案	
	Diffie-Hellman 迪菲-赫尔曼密钥交换
## 加密方法
	自定义加密 AES DES RC4

```
self encrypt 26 26 460ns
self decrypt 26 26 441ns
aes encrypt 26 26 370ns
aes decrypt 26 26 107ns
des encrypt 26 26 888ns
des decrypt 26 26 307ns
rc4 encrypt 26 26 249ns
rc4 decrypt 26 26 161ns

self encrypt 32 32 453ns
self decrypt 32 32 486ns
aes encrypt 32 32 70ns
aes decrypt 32 32 103ns
des encrypt 32 32 267ns
des decrypt 32 32 244ns
rc4 encrypt 32 32 146ns
rc4 decrypt 32 32 176ns

self encrypt 24 24 345ns
self decrypt 24 24 373ns
aes encrypt 24 24 79ns
aes decrypt 24 24 96ns
des encrypt 24 24 231ns
des decrypt 24 24 240ns
rc4 encrypt 24 24 148ns
rc4 decrypt 24 24 128ns

self encrypt 255 255 3.386µs
self decrypt 255 255 3.441µs
aes encrypt 255 255 369ns
aes decrypt 255 255 235ns
des encrypt 255 255 530ns
des decrypt 255 255 453ns
rc4 encrypt 255 255 871ns
rc4 decrypt 255 255 760ns

self encrypt 1025 1025 22.232µs
self decrypt 1025 1025 17.495µs
aes encrypt 1025 1025 286ns
aes decrypt 1025 1025 120ns
des encrypt 1025 1025 466ns
des decrypt 1025 1025 298ns
rc4 encrypt 1025 1025 2.352µs
rc4 decrypt 1025 1025 2.334µs

self encrypt 2045 2045 36.777µs
self decrypt 2045 2045 32.647µs
aes encrypt 2045 2045 241ns
aes decrypt 2045 2045 218ns
des encrypt 2045 2045 717ns
des decrypt 2045 2045 441ns
rc4 encrypt 2045 2045 4.645µs
rc4 decrypt 2045 2045 4.216µs

self encrypt 4045 4045 82.013µs
self decrypt 4045 4045 104.377µs
aes encrypt 4045 4045 718ns
aes decrypt 4045 4045 488ns
des encrypt 4045 4045 2.135µs
des decrypt 4045 4045 517ns
rc4 encrypt 4045 4045 9.03µs
rc4 decrypt 4045 4045 8.412µs
```
