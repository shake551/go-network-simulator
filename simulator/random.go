package simulator

import (
	"math"
	"math/rand"
	"time"
)

func RandomMillisecond(x float64) int {
	rand.Seed(time.Now().UnixNano())

	expRandom := math.Log(1-rand.Float64()) * -1 / x
	return int(expRandom * 1000)
}
