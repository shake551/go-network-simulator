package main

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
	"time"
)

func main() {
	nowTime := time.Now()
	finishTime := nowTime.Add(time.Second * 5)

	s := simulator.NewSystem(0.5, 0.6, 1000, nowTime, finishTime, 1000)
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
