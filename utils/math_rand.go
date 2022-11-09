package utils

import (
	"math/rand"
	"time"
)

func MathRand() float64 {
	rand.Seed(time.Now().UnixNano())

	return rand.Float64()
}
