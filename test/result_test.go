package test

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
	"testing"
)

func TestResult(t *testing.T) {
	totalPacketCount := 0.0
	totalPacketStayTime := 0.0
	totalPacketLossRate := 0.0

	for i := 0; i < 10; i++ {
		startTime := 0.0
		finishTime := 10000.0
		s := simulator.NewSystem(0.95, 1, startTime, finishTime, 50)
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
				stayTime := (*s.PacketStatistics).GetAverageOfPacketStayTime()
				packetLossRate := (*s.PacketStatistics).GetPacketLossRate()

				totalPacketCount += packetCount
				totalPacketStayTime += stayTime
				totalPacketLossRate += packetLossRate

				break
			}
		}

		fmt.Println()

		s.ShowTheoreticalValues()

		fmt.Println("-------------- show statistics ----------------")
		fmt.Printf("average of packet conunt: %f\naverage of packet stay time: %f\npacket loss rate: %f\n", totalPacketCount/10, totalPacketStayTime/10, totalPacketLossRate/10)
	}
}
