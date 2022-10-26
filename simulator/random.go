package simulator

import (
	"crypto/rand"
	"math"
	"math/big"
)

func RandomMillisecond(x float64) float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}

	expRandom := math.Log(1-float64(n.Int64())/100) * -1 / x
	return expRandom
}
