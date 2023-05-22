package bls_package

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"

	"github.com/herumi/bls-eth-go-binary/bls"
)

func BlsBenchmark() {
	bls.Init(bls.BLS12_381)
	bls.SetETHmode(bls.EthModeDraft07)
	//创建num个签名
	num := 10000
	//创建私钥切片
	//创建公钥切片

	var privs []*bls.SecretKey = make([]*bls.SecretKey, num)
	var pubs []bls.PublicKey = make([]bls.PublicKey, num)
	var sigs []bls.Sign = make([]bls.Sign, num)
	start := time.Now() // 获取当前时间
	var sec bls.SecretKey
	for i := 0; i < num; i++ {
		sec.SetByCSPRNG()
		msg := []byte("hello world")
		privs[i] = &sec
		pub := sec.GetPublicKey()
		pubs[i] = *pub
		sig := sec.SignByte(msg)
		sigs[i] = *sig
	}

	elapsed := time.Now().Sub(start)
	fmt.Println("生成"+strconv.Itoa(num)+"个签名完成耗时：", elapsed)
	//聚合签名

	var aggSig bls.Sign
	//计算时间
	start = time.Now() // 获取当前时间
	//遍历签名切片，将签名聚合
	for i := 0; i < num; i++ {
		aggSig.Add(&sigs[i])
	}
	elapsed = time.Now().Sub(start)
	fmt.Println("聚合"+strconv.Itoa(num)+"个签名完成耗时：", elapsed)

	//遍历公钥切片，将公钥聚合
	//计算时间
	start = time.Now() // 获取当前时间

	var aggPub bls.PublicKey
	for i := 0; i < num; i++ {
		aggPub.Add(&pubs[i])
	}
	elapsed = time.Now().Sub(start)
	fmt.Println("聚合"+strconv.Itoa(num)+"个公钥完成耗时：", elapsed)
	//验证聚合签名
	//输出聚合后的签名大小
	fmt.Println("聚合后的签名大小：", unsafe.Sizeof(&aggSig))
	//计算时间
	start = time.Now() // 获取当前时间
	result := aggSig.VerifyByte(&aggPub, []byte("hello world"))
	elapsed = time.Now().Sub(start)
	fmt.Println("验证"+strconv.Itoa(num)+"个签名完成耗时：", elapsed)
	print(result)
}
