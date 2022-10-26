package main

import (
	"fmt"
	"github.com/shake551/go-network-simulator/simulator"
	"github.com/shake551/go-network-simulator/utils"
)

func exeSimulate(packetRate float64) {
	startTime := 0.0
	finishTime := 10000.0

	totalPacketCount := 0.0
	totalPacketStayTime := 0.0
	totalPacketLossRate := 0.0

	repeatCount := 1000

	gs := utils.NewGSpread(packetRate)
	for i := 0; i < repeatCount; i++ {
		s := simulator.NewSystem(packetRate, 1, startTime, finishTime, 50)
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

				gs.AppendNewStatistics(packetCount, stayTime, packetLossRate)

				totalPacketCount += packetCount
				totalPacketStayTime += stayTime
				totalPacketLossRate += packetLossRate

				s.ShowTheoreticalValues()

				break
			}
		}
	}
	fmt.Println()
	fmt.Printf("packetRate: %f\n", packetRate)

	fmt.Println("-------------- show statistics ----------------")
	fmt.Printf("average of packet conunt: %f\naverage of packet stay time: %f\npacket loss rate: %f\n",
		totalPacketCount/float64(repeatCount), totalPacketStayTime/float64(repeatCount), totalPacketLossRate/float64(repeatCount))
	gs.Insert()
}

func main() {
	packetRates := [7]float64{0.7, 0.75, 0.8, 0.85, 0.9, 0.95, 1}
	for _, packetRate := range packetRates {
		exeSimulate(packetRate)
		fmt.Println(packetRate)
	}

	fmt.Println("finish")
}
