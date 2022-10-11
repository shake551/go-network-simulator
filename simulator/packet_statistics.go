package simulator

type PacketStatistics struct {
	TotalCount    int
	PacketLoss    int
	TotalStayTime float64
}

func (p PacketStatistics) GetPacketLossRate() float64 {
	return (float64(p.PacketLoss) / float64(p.TotalCount)) * 100
}

func (p PacketStatistics) GetAverageOfPacketStayTime(totalTime float64) float64 {
	return p.TotalStayTime / totalTime
}
