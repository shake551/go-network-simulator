package main

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
)

func main() {
	startTime := 0.0
	finishTime := 100.0
	s := simulator.NewSystem(0.5, 0.6, 1000, startTime, finishTime, 1000)
	s.Init()

	for true {
		keep, err := s.Simulate()
		if err != nil {
			fmt.Println(err)
		}

		if !keep {
			fmt.Println("finish time")
			return
		}
	}
}
