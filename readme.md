Ether Utils
===========

Generic utilities for messing around with ethereum because I get annoyed at trying to find my
most recent copy of each of these....

StrToDecimals
---

* convert a string (possibly with a decimal point) to the atomic unit based on the number of decimals expected

```
func StrToDecimals(value string, decimals int64) (vInEther *big.Int, ok bool)
```

StrToEther
---

* Convert a string (possibly with a decimal point) to Wei

```
func StrToEther(value string) (vInEther *big.Int, ok bool)
```

EtherToStr
---

* Convert a value (in wei) to an ether string

```
func EtherToStr(bbal *big.Int) string
```

CoinToString
---

* Convert a value (in uCoins) to a Coin string

```
func CoinToStr(bbal *big.Int, numDec int) string
```

StrToDecimals
-------------

* convert a string like "10.32" into the base units given the number of decimal places

