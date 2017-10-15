package etherUtils

import "math/big"

func bi(n int) *big.Int {
	return new(big.Int).SetUint64(uint64(n))
}

func oneEther() *big.Int {
	return new(big.Int).SetUint64(1000000000000000000)
}
