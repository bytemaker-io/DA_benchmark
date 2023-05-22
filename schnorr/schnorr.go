package schnorr

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/hbakhtiyor/schnorr"
)

func SchnorrBenchmark() {
	Curve := btcec.S256()
	//创建num个私钥
	num := 10000
	//创建一个私钥切片

	privateKeys := make([]btcec.PrivateKey, num)
	start := time.Now() // 获取当前时间
	for i := 0; i < num; i++ {
		p, _ := btcec.NewPrivateKey(Curve)
		privateKeys[i] = *p
	}
	elapsed := time.Since(start)
	fmt.Println("创建"+strconv.Itoa(num)+"个私钥耗时:", elapsed)

	var (
		publicKey [33]byte
		message   [32]byte
	)

	msg, _ := hex.DecodeString("243F6A8885A308D313198A2E03707344A4093822299F31D0082EFA98EC4E6C89")
	copy(message[:], msg)

	//创建私钥D切片
	privateKeysD := make([]*big.Int, num)
	for i := 0; i < num; i++ {
		privateKeysD[i] = privateKeys[i].D
	}
	//使用聚合私钥创建签名
	start = time.Now() // 获取当前时间
	signature, _ := schnorr.AggregateSignatures(privateKeysD, message)
	fmt.Println("生成签名并聚合耗时:", time.Since(start))
	//循环调用Curve.Add,聚合所有公钥
	start = time.Now() // 获取当前时间
	Px, Py := Curve.Add(privateKeys[0].PublicKey.X, privateKeys[0].PublicKey.Y, privateKeys[1].PublicKey.X, privateKeys[1].PublicKey.Y)
	for i := 2; i < num; i++ {
		Px, Py = Curve.Add(Px, Py, privateKeys[i].PublicKey.X, privateKeys[i].PublicKey.Y)
	}
	copy(publicKey[:], schnorr.Marshal(Curve, Px, Py))
	fmt.Println("聚合公钥耗时:", time.Since(start))
	start = time.Now() // 获取当前时间
	if result, err := schnorr.Verify(publicKey, message, signature); err != nil {
		fmt.Printf("The signature verification failed: %v\n", err)
	} else if result {
		fmt.Println("The signature is valid.")
	}
	fmt.Println("验证耗时:", time.Since(start))

}
