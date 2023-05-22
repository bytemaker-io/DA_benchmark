package schnorr

import (
	"fmt"

	"github.com/alinush/bls-go-binding"
)

func SchnorrBenchmark() {
	const numSigs = 100000

	// 生成私钥和公钥
	var privKeys []*bls.SecretKey
	var pubKeys []*bls.PublicKey
	for i := 0; i < numSigs; i++ {
		privKey := new(bls.SecretKey)

		privKeys = append(privKeys, privKey)

		pubKey := privKey.GetPublicKey()
		pubKeys = append(pubKeys, pubKey)
	}

	// 签名消息
	msg := "Hello World!"
	var sigs []*bls.Sign
	for i := 0; i < numSigs; i++ {
		sig := privKeys[i].Sign(msg)
		sigs = append(sigs, sig)
	}

	// 聚合签名
	var aggSig *bls.Sign
	for i := 0; i < numSigs; i++ {
		aggSig.Add(sigs[i])
	}

	// 聚合公钥
	var aggPubKey *bls.PublicKey
	for i := 0; i < numSigs; i++ {
		aggPubKey.Add(pubKeys[i])
	}

	// 验证聚合后的签名
	if aggSig.Verify(aggPubKey, msg) {
		fmt.Println("Aggregated signature is valid")
	} else {
		fmt.Println("Aggregated signature is invalid")
	}
}
