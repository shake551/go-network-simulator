package test

import (
	"github.com/shake551/go-network-simulator/simulator"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	nowTime := time.Now()
	finishTime := nowTime.Add(time.Minute * 5)

	s := simulator.NewSystem(0.5, 0.6, 1000, nowTime, finishTime, 1000)

	s.Init()

	if s.PacketRate != 0.5 {
		t.Errorf("the packetRate should be 0.5, but got %f", s.PacketRate)
	}

	if s.SystemCapacity != 1000 {
		t.Errorf("the system capacity should be 1000, but got %d", s.SystemCapacity)
	}

	if s.ServiceRate != 0.6 {
		t.Errorf("the serviceRate should be 0.6, but got %f", s.ServiceRate)
	}

	if s.StartTime != nowTime.UnixMicro() {
		t.Errorf("the start time should be %d, but got %d", nowTime.UnixMicro(), s.StartTime)
	}

	if s.FinishTime != finishTime.UnixMicro() {
		t.Errorf("the finish time should be %d, but got %d", finishTime.UnixMicro(), s.FinishTime)
	}

	if s.NowTime != nowTime.UnixMicro() {
		t.Errorf("the nowTime should be %d, but got %d", nowTime.UnixMicro(), s.NowTime)
	}

	if len(*s.EventTable) != 2 {
		t.Errorf("the length of event table should be 2, but got %d", len(*s.EventTable))
	}
}
