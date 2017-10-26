package etherUtils

import (
	"fmt"
	"math/big"
	"strings"
)

// StrToDecimals converts 1.0 to 1, 10, 100 etc depending on the decimals parameter
func StrToDecimals(value string, decimals int64) (vInEther *big.Int, ok bool) {
	v, ok := new(big.Int).SetString(value, 10)
	//fmt.Println(v, ok)
	powerInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)
	if ok {
		vInEther = new(big.Int).Mul(v, powerInt)
		return
	}
	strA := strings.Split(value, ".")
	if len(strA) != 2 {
		strA = append(strA, "0")
	}
	v, ok = new(big.Int).SetString(strA[0], 10)
	if !ok {
		fmt.Println("1")
		return
	}
	vInWholeEther := new(big.Int).Mul(v, powerInt)
	v2, ok := new(big.Int).SetString(strA[1], 10)
	if !ok {
		fmt.Println("2")
		return
	}
	pwr := new(big.Int).Exp(
		bi(10),
		bi(int(decimals-int64(len(strA[1])))),
		nil)
	vInPartEther := new(big.Int).Mul(v2, pwr)
	vInEther = new(big.Int).Add(vInWholeEther, vInPartEther)
	//fmt.Println(strA[0], strA[1])
	return
}

// StrToEther - "1.0" -> 1 + 10 zeroes
func StrToEther(value string) (vInEther *big.Int, ok bool) {
	return StrToDecimals(value, 18)
}

// EtherToStr - given a number of wei, express it as a decimal amount of ether
func EtherToStr(bbal *big.Int) string {
	return CoinToStr(bbal, 18)
}

// CoinToStr - convert coin from uCoins to fractional coin strings
func CoinToStr(bbal *big.Int, numDec int) string {
	fStr := fmt.Sprintf("%%0%dd", numDec)
	bb := fmt.Sprintf(fStr, bbal)
	if len(bb) == numDec {
		bb = "0." + bb
	} else {
		l := len(bb)
		bb = fmt.Sprintf("%s.%s", bb[:l-numDec], bb[l-numDec:])
	}
	return bb
}
