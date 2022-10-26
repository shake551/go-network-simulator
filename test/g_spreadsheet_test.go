package test

import (
	"github.com/shake551/go-network-simulator/utils"
	"testing"
)

func TestConnection(t *testing.T) {
	_ = utils.NewGSpread(0.7)
}

func TestCreateNewSheet(t *testing.T) {
	gs := utils.NewGSpread(0.7)
	gs.CreateNewSheet("test create new sheet")
}
