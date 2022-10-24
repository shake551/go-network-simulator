package main

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
)

func main() {
	startTime := 0.0
	finishTime := 10000.0
	s := simulator.NewSystem(0.7, 1, startTime, finishTime, 50)
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
			fmt.Println()

			s.ShowTheoreticalValues()

			fmt.Println("-------------- show statistics ----------------")
			fmt.Printf("average of packet conunt: %f\naverage of packet stay time: %f\npacket loss rate: %f\n", packetCount, (*s.PacketStatistics).GetAverageOfPacketStayTime(), (*s.PacketStatistics).GetPacketLossRate())
			return
		}
	}
}
