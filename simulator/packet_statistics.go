package simulator

type PacketStatistics struct {
	TotalCount    int
	PacketLoss    int
	TotalStayTime float64
}

func (p PacketStatistics) GetPacketLossRate() float64 {
	return float64(p.PacketLoss) / float64(p.TotalCount)
}

func (p PacketStatistics) GetAverageOfPacketCount(totalTime float64) float64 {
	return p.TotalStayTime / totalTime
}

func (p PacketStatistics) GetAverageOfPacketStayTime() float64 {
	return p.TotalStayTime / float64(p.TotalCount)
}
