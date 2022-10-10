package simulator

import (
	"fmt"
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
	NowEvent       *Event
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
		NowEvent:       &Event{Type: "start", Time: startTime.UnixMicro()},
		EventTable:     &[]Event{},
		EventQueue:     &EventQueue{MaxSize: maxSize},
		IsProcess:      utils.Bool(false),
	}
}

func (s System) Init() {
	*s.EventTable = append(*s.EventTable, Event{
		Type: "eventFinish",
		Time: s.FinishTime,
	})

	s.AppendEvent("start")

	s.MakeProcess()
}

func (s System) AppendEvent(eventType string) {
	nowTime := time.UnixMicro(s.NowEvent.Time)
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

func (s System) MoveToNextEvent() error {
	s.SortEventTableByTime()

	if len(*s.EventTable) == 0 {
		return fmt.Errorf("event table is empty")
	}

	*s.NowEvent = (*s.EventTable)[0]
	*s.EventTable = (*s.EventTable)[1:]

	return nil
}

func (s System) EventStart() {
	s.AppendEvent("start")

	if !*s.IsProcess {
		s.MakeProcess()
		return
	}

	if !(*s.EventQueue).IsFull() {
		(*s.EventQueue).Enqueue(*s.NowEvent)
	}
}

func (s System) EventFinish() {
	queueEvent := (*s.EventQueue).Dequeue()
	if queueEvent.Type == "" {
		fmt.Println("there is no queued event")
		return
	}

	s.AppendEvent("finish")
	s.MakeProcess()
}

func (s System) GetPacketTime() int {
	return RandomMillisecond(s.PacketRate)
}

func (s System) GetServiceTime() int {
	return RandomMillisecond(s.ServiceRate)
}

func (s System) MakeProcess() {
	*s.IsProcess = true
	s.AppendEvent("finish")
}

func (s System) UnProcess() {
	*s.IsProcess = false
}
