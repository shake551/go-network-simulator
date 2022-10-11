package simulator

import (
	"math"
	"math/rand"
	"time"
)

func RandomMillisecond(x float64) float64 {
	rand.Seed(time.Now().UnixNano())

	expRandom := math.Log(1-rand.Float64()) * -1 / x
	return expRandom
}
