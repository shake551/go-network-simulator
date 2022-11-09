package main

import (
	"github.com/shake551/go-network-simulator/utils"
	"testing"
)

func TestRandomNumberDistribution(t *testing.T) {
	repeatCount := 1000

	gs := utils.NewGSpread(-0.11)
	for i := 0; i < repeatCount; i++ {
		gs.AppendNewStatistics(utils.CryptoRand(), utils.MathRand(), 0)
	}

	gs.Insert()
}
