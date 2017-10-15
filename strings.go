package etherUtils

import (
	"fmt"
	"math/big"
	"strings"
)

// EtherStrToWei converts 1.0 to 1000000000000000000000
//                        0.1 to  100000000000000000000
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

func StrToEther(value string) (vInEther *big.Int, ok bool) {
	return StrToDecimals(value, 18)
}

func EtherToStr(bbal *big.Int) string {
	bb := fmt.Sprintf("%018d", bbal)
	if len(bb) == 18 {
		bb = "0." + bb
	} else {
		l := len(bb)
		bb = fmt.Sprintf("%s.%s", bb[:l-18], bb[l-18:])
	}
	return bb
}
