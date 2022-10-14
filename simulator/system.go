package simulator

import (
	"fmt"
	"github.com/shake551/go-network-simulator/utils"
	"sort"
)

type System struct {
	PacketRate       float64
	ServiceRate      float64
	StartTime        float64
	FinishTime       float64
	NowEvent         *Event
	EventTable       *[]Event
	EventQueue       *EventQueue
	IsProcess        *bool
	PacketStatistics *PacketStatistics
}

func NewSystem(packetRate float64, serviceRate float64, startTime float64, finishTime float64, maxSize int) *System {
	return &System{
		PacketRate:       packetRate,
		ServiceRate:      serviceRate,
		StartTime:        startTime,
		FinishTime:       finishTime,
		NowEvent:         &Event{Type: "start", Time: startTime},
		EventTable:       &[]Event{},
		EventQueue:       &EventQueue{MaxSize: maxSize - 1},
		IsProcess:        utils.Bool(false),
		PacketStatistics: &PacketStatistics{TotalCount: 0, PacketLoss: 0, TotalStayTime: 0},
	}
}

func (s System) Init() {
	*s.EventTable = append(*s.EventTable, Event{
		Type: "eventFinish",
		Time: s.FinishTime,
	})

	s.AppendStartEvent()

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

func (s System) AppendStartEvent() {
	eventTime := s.NowEvent.Time + s.GetPacketTime()

	*s.EventTable = append(*s.EventTable, Event{
		Type: "start",
		Time: eventTime,
	})
}

func (s System) AppendFinishEvent() {
	eventTime := s.NowEvent.Time + s.GetServiceTime()

	*s.EventTable = append(*s.EventTable, Event{
		Type:          "finish",
		Time:          eventTime,
		InServiceTime: s.NowEvent.Time,
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
	s.AppendStartEvent()

	(*s.PacketStatistics).TotalCount++

	if !*s.IsProcess {
		s.MakeProcess()
		return
	}

	err := (*s.EventQueue).Enqueue(*s.NowEvent)
	if err != nil {
		(*s.PacketStatistics).PacketLoss++
	}
}

func (s System) EventFinish() {
	(*s.PacketStatistics).TotalStayTime += (*s.NowEvent).Time - (*s.NowEvent).InServiceTime

	queueEvent := (*s.EventQueue).Dequeue()
	if queueEvent.Type == "" {
		fmt.Println("there is no queued event")
		s.UnProcess()
		return
	}

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
	s.AppendFinishEvent()
}

func (s System) UnProcess() {
	*s.IsProcess = false
}
