package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func EcdsaBenchMark() {
	curve := elliptic.P256()
	const num = 100000
	// 生成一万个私钥ecdsa签名私钥
	var privateKeySlice []*ecdsa.PrivateKey = make([]*ecdsa.PrivateKey, num, num)

	// 生成一万个私钥
	for i := 0; i < num; i++ {
		// 生成私钥
		privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			panic(err)
		}
		privateKeySlice[i] = privateKey
	}

	var rSlice []*big.Int = make([]*big.Int, num, num)
	var sSlice []*big.Int = make([]*big.Int, num, num)
	// 生成一万个签名
	start := time.Now() // 获取当前时间
	for i := 0; i < num; i++ {
		// 生成签名
		r, s, err := ecdsa.Sign(rand.Reader, privateKeySlice[i], []byte("hello world"))
		if err != nil {
			panic(err)
		}
		//将r和s存到切片中
		rSlice[i] = r
		sSlice[i] = s
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("生成"+strconv.Itoa(num)+"个签名完成耗时：", elapsed)

	//验证一万个签名
	start = time.Now() // 获取当前时间
	for i := 0; i < num; i++ {
		// 验证签名
		if !ecdsa.Verify(&privateKeySlice[i].PublicKey, []byte("hello world"), rSlice[i], sSlice[i]) {
			panic("验证签名失败")
		}
	}
	elapsed = time.Now().Sub(start)
	fmt.Println("验证"+strconv.Itoa(num)+"个签名完成耗时：", elapsed)
}
