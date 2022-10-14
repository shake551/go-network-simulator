package main

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
)

func main() {
	startTime := 0.0
	finishTime := 1000.0
	s := simulator.NewSystem(1, 0.1, startTime, finishTime, 100)
	s.Init()

	for true {
		keep, err := s.Simulate()
		if err != nil {
			fmt.Println(err)
		}

		if !keep {
			fmt.Println("finish time")

			s.AddStayTimeOfEventsInQueue()

			packetCount := (*s.PacketStatistics).GetAverageOfPacketCount(s.FinishTime - s.StartTime)
			fmt.Printf("average of packet conunt: %f, average of packet stay time: %f, packet loss rate: %f", packetCount, (*s.PacketStatistics).GetAverageOfPacketStayTime(), (*s.PacketStatistics).GetPacketLossRate())
			return
		}
	}
}
