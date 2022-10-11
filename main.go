package main

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
)

func main() {
	startTime := 0.0
	finishTime := 100.0
	s := simulator.NewSystem(1, 0.1, startTime, finishTime, 10)
	s.Init()

	for true {
		keep, err := s.Simulate()
		if err != nil {
			fmt.Println(err)
		}

		if !keep {
			fmt.Println("finish time")

			packetStayTime := (*s.PacketStatistics).GetAverageOfPacketStayTime(s.FinishTime - s.StartTime)
			fmt.Printf("average of packet stay time: %f, packet loss rate: %f", packetStayTime, (*s.PacketStatistics).GetPacketLossRate())
			return
		}
	}
}
