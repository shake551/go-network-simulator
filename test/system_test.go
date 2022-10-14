package test

import (
	"github.com/shake551/go-network-simulator/simulator"
	"testing"
)

func TestInit(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()

	if s.PacketRate != 0.5 {
		t.Errorf("the packetRate should be 0.5, but got %f", s.PacketRate)
	}

	if s.ServiceRate != 0.6 {
		t.Errorf("the serviceRate should be 0.6, but got %f", s.ServiceRate)
	}

	if s.StartTime != startTime {
		t.Errorf("the start time should be %f, but got %f", startTime, s.StartTime)
	}

	if s.FinishTime != finishTime {
		t.Errorf("the finish time should be %f, but got %f", finishTime, s.FinishTime)
	}

	if s.NowEvent.Time != startTime {
		t.Errorf("the nowEventTime should be %f, but got %f", startTime, s.NowEvent.Time)
	}

	if len(*s.EventTable) != 3 {
		t.Errorf("the length of event table should be 3, but got %d", len(*s.EventTable))
	}
}

func TestAppendEvent(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()

	s.AppendStartEvent()

	targetEvent := (*s.EventTable)[len(*s.EventTable)-1]

	if targetEvent.Type != "start" {
		t.Errorf("the event type shold be start, but got %s", targetEvent.Type)
	}

	if targetEvent.Time <= s.NowEvent.Time {
		t.Errorf("the event time should be bigger than nowTime")
	}
}

func TestSortEventTable(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()
	s.AppendStartEvent()
	s.SortEventTableByTime()

	finishEvent := (*s.EventTable)[len(*s.EventTable)-1]

	if finishEvent.Type != "eventFinish" {
		t.Errorf("the event type should be eventFinish, but got %s", finishEvent.Type)
	}

	if finishEvent.Time != finishTime {
		t.Errorf("the finish time should be %f, but got %f", finishEvent.Time, finishTime)
	}
}

func TestIsProcess(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()
	if !*s.IsProcess {
		t.Errorf("the isProcess should be true, but got false")
	}

	s.UnProcess()
	if *s.IsProcess {
		t.Errorf("the isProcess should be false, but got true")
	}

	s.MakeProcess()
	if !*s.IsProcess {
		t.Errorf("the isProcess should be true, but got false")
	}
}

func TestMoveToNextEvent(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()

	err := s.MoveToNextEvent()
	if err != nil {
		t.Errorf("cannot move to next event")
	}

	if (*s.NowEvent).Time == startTime {
		t.Errorf("nowEventTime should be more than %f, but got %f", (*s.NowEvent).Time, startTime)
	}

	if (*s.NowEvent).Type == (*s.EventTable)[0].Type {
		t.Errorf("the NowEvent type should be different from next event type, but got %s = %s", (*s.NowEvent).Type, (*s.EventTable)[0].Type)
	}

	*s.EventTable = []simulator.Event{}

	err = s.MoveToNextEvent()
	if err == nil {
		t.Errorf("the function MoveToNextEvent should return error, but got nil")
	}
}

func TestEventStart(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()
	s.EventStart()

	if len((*s.EventQueue).Data) != 1 {
		t.Errorf("the length of event queue should be 1, but got %d", len((*s.EventQueue).Data))
	}

	*s.IsProcess = false
	s.EventStart()

	if !*s.IsProcess {
		t.Errorf("the simulator should be pricessing, but not")
	}

	// testing do not update event queue
	if len((*s.EventQueue).Data) != 1 {
		t.Errorf("the length of event queue should be 1, but got %d", len((*s.EventQueue).Data))
	}
}

func TestEventFinish(t *testing.T) {
	startTime := 0.0
	finishTime := 10.0
	s := simulator.NewSystem(0.5, 0.6, startTime, finishTime, 1000)

	s.Init()

	for (*s.NowEvent).Type != "finish" {
		s.SortEventTableByTime()
		s.EventStart()
		err := s.MoveToNextEvent()
		if err != nil {
			t.Errorf("cannot move to next event")
		}
	}

	s.SortEventTableByTime()
	if (*s.NowEvent).Type != "finish" {
		t.Errorf("nowEvent type should be finish, but got %s", (*s.NowEvent).Type)
	}
	s.EventFinish()

	if !*s.IsProcess {
		t.Errorf("the simulator should be processing")
	}
}
