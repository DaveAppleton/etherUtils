package etherUtils

import "math/big"

func bi(n int) *big.Int {
	return new(big.Int).SetUint64(uint64(n))
}

// OneEther as a big Int
func OneEther() *big.Int {
	return new(big.Int).SetUint64(1000000000000000000)
}

// PointOneEther as a big int
func PointOneEther() *big.Int {
	return new(big.Int).SetUint64(100000000000000000)
}
