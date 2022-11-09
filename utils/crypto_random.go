package utils

import (
	"crypto/rand"
	"math/big"
)

func CryptoRand() float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}

	return float64(n.Int64()) / 100
}
