## 概要
    加密算法的分类 
        1. 不可逆加密
        2. 可逆加密
		   可逆加密又分为对称加密及非对称加密

    这个示例主要是演示及对比可逆加密,希望看完后对项目的加密选择有一定的帮助
	
## 可逆加密对比
### 加密方法
    对称加密方法
	    aes,des,rc4
	非对称加密方法
	    rsa
	AES DES RC4

### 非对称加密key的交换
	OAEP  最佳非对称加密填充（OAEP）是一个通常和RSA一起使用的填充方案	
	Diffie-Hellman 迪菲-赫尔曼密钥交换

#### 64字节长度的加密性能测试结果

```
64
goos: darwin
goarch: amd64
Benchmark_aes-4         20000000                62.5 ns/op
Benchmark_des-4          5000000               304 ns/op
Benchmark_rc4-4         10000000               220 ns/op
```
