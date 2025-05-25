package myWallet

import (
	"math/big"
)

var decimalsMap map[uint32]big.Float

func init() {
	decimalsMap = make(map[uint32]big.Float)
	for i := 1; i <= 30; i++ {
		dec := *NewFloat256ByInt64(10)
		d10 := *NewFloat256ByInt64(10)
		for j := 1; j < i; j++ {
			dec = Float256Mul(dec, d10)
		}
		decimalsMap[uint32(i)] = dec
	}

}

// OriginalToFace 区块原始金额转面值
func OriginalToFace(amount string, decimals uint32) big.Float {
	if 0 == decimals {
		return *NewFloat256ByStringMust(amount)
	}
	dec := decimalsMap[decimals]
	a := NewFloat256ByStringMust(amount)
	s := Float256Quo(*a, dec)
	return s
}

// FaceToOriginal 面额转区块原始金额
func FaceToOriginal(amount string, decimals uint32) big.Float {
	if 0 == decimals {
		return *NewFloat256ByStringMust(amount)
	}
	dec := decimalsMap[decimals]
	a := NewFloat256ByStringMust(amount)
	s := Float256Mul(*a, dec)
	return s
}
