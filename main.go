package main

import (
	"DA_benchmark/bls_package"
	"DA_benchmark/ecdsa"
	"DA_benchmark/schnorr"
)

func main() {

	bls_package.BlsBenchmark()
	schnorr.SchnorrBenchmark()
	ecdsa.EcdsaBenchMark()
}
