package simulator

import (
	"fmt"
	"github.com/shake551/go-network-simulator/utils"
	"sort"
)

type System struct {
	PacketRate     float64
	ServiceRate    float64
	SystemCapacity int64
	StartTime      float64
	FinishTime     float64
	NowEvent       *Event
	EventTable     *[]Event
	EventQueue     *EventQueue
	IsProcess      *bool
}

func NewSystem(packetRate float64, serviceRate float64, systemCapacity int64, startTime float64, finishTime float64, maxSize int) *System {
	return &System{
		PacketRate:     packetRate,
		ServiceRate:    serviceRate,
		SystemCapacity: systemCapacity,
		StartTime:      startTime,
		FinishTime:     finishTime,
		NowEvent:       &Event{Type: "start", Time: startTime},
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

func (s System) Simulate() (bool, error) {
	err := s.MoveToNextEvent()
	if err != nil {
		return false, err
	}

	fmt.Printf("simulate %f, %s\n", (*s.NowEvent).Time, (*s.NowEvent).Type)

	switch (*s.NowEvent).Type {
	case "start":
		s.EventStart()
	case "finish":
		s.EventFinish()
	case "eventFinish":
		return false, nil
	default:
		return false, fmt.Errorf("unexpected input: %s", (*s.NowEvent).Type)
	}

	return true, nil
}

func (s System) AppendEvent(eventType string) {
	eventTime := s.NowEvent.Time + s.GetPacketTime()

	*s.EventTable = append(*s.EventTable, Event{
		Type: eventType,
		Time: eventTime,
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
		s.UnProcess()
		return
	}

	s.AppendEvent("finish")
	s.MakeProcess()
}

func (s System) GetPacketTime() float64 {
	return RandomMillisecond(s.PacketRate)
}

func (s System) GetServiceTime() float64 {
	return RandomMillisecond(s.ServiceRate)
}

func (s System) MakeProcess() {
	*s.IsProcess = true
	s.AppendEvent("finish")
}

func (s System) UnProcess() {
	*s.IsProcess = false
}
