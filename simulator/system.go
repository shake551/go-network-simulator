package simulator

import (
	"time"
)

type System struct {
	PacketRate     float64
	ServiceRate    float64
	SystemCapacity int64
	StartTime      int64
	FinishTime     int64
	NowTime        int64
	EventTable     *[]Event
	EventQueue     *EventQueue
}

func NewSystem(packetRate float64, serviceRate float64, systemCapacity int64, startTime time.Time, finishTime time.Time, maxSize int) *System {
	return &System{
		PacketRate:     packetRate,
		ServiceRate:    serviceRate,
		SystemCapacity: systemCapacity,
		StartTime:      startTime.UnixMicro(),
		FinishTime:     finishTime.UnixMicro(),
		NowTime:        startTime.UnixMicro(),
		EventTable:     &[]Event{{Type: "start", Time: startTime.UnixMicro()}},
		EventQueue:     &EventQueue{MaxSize: maxSize},
	}
}

func (s System) Init() {
	*s.EventTable = append(*s.EventTable, Event{
		Type: "finish",
		Time: s.FinishTime,
	})
}
