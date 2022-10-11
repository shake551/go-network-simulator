package test

import (
	"github.com/shake551/go-network-simulator/simulator"
	"testing"
)

func TestQueue(t *testing.T) {
	q := &simulator.EventQueue{MaxSize: 5}

	if !q.IsEmpty() {
		t.Error("the queue should be empty, but it's not.")
	}

	for i := 0; i < 5; i++ {
		e := simulator.Event{Type: "start", Time: float64(i)}
		q.Enqueue(e)
	}

	if !q.IsFull() {
		t.Error("the queue should be full, but it's not.")
	}

	e := q.Dequeue()
	if e.Type != "start" {
		t.Errorf("the dequeue Type should be start, but got %s", e.Type)
	}

	if len(q.Data) != 4 {
		t.Errorf("the queued data length should be 4, but got %d", len(q.Data))
	}
}
