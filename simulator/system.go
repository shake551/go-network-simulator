package simulator

import (
	"github.com/shake551/go-network-simulator/utils"
	"sort"
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
	IsProcess      *bool
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
		IsProcess:      utils.Bool(false),
	}
}

func (s System) Init() {
	*s.EventTable = append(*s.EventTable, Event{
		Type: "finish",
		Time: s.FinishTime,
	})
}

func (s System) AppendEvent(eventType string) {
	nowTime := time.UnixMicro(s.NowTime)
	durationMillisecond := time.Duration(s.GetPacketTime())

	eventTime := nowTime.Add(time.Millisecond * durationMillisecond)

	*s.EventTable = append(*s.EventTable, Event{
		Type: eventType,
		Time: eventTime.UnixMicro(),
	})
}

func (s System) SortEventTableByTime() {
	sort.Slice(*s.EventTable, func(i, j int) bool { return (*s.EventTable)[i].Time < (*s.EventTable)[j].Time })
}

func (s System) GetPacketTime() int {
	return RandomMillisecond(s.PacketRate)
}

func (s System) GetServiceTime() int {
	return RandomMillisecond(s.ServiceRate)
}

func (s System) MakeProcess() {
	*s.IsProcess = true
}

func (s System) UnProcess() {
	*s.IsProcess = false
}
