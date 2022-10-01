package simulator

import (
	"math"
	"math/rand"
	"time"
)

func ExpRand(x float64) float64 {
	rand.Seed(time.Now().UnixNano())

	return math.Log(1-rand.Float64()) * -1 / x
}
