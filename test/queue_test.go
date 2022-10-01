package test

import (
	"testing"
	"time"

	"github.com/shake551/go-network-simulator/simulator"
)

func TestQueue(t *testing.T) {
	q := &simulator.EventQueue{MaxSize: 5}

	if !q.IsEmpty() {
		t.Error("the queue should be empty, but it's not.")
	}

	for i := 0; i < 5; i++ {
		e := simulator.Event{ID: i, Type: "start", Time: time.Now().Unix()}
		q.Enqueue(e)
	}

	if !q.IsFull() {
		t.Error("the queue should be full, but it's not.")
	}

	e := q.Dequeue()
	if e.ID != 0 {
		t.Errorf("the dequeue ID should be 0, but got %d", e.ID)
	}

	if len(q.Data) != 4 {
		t.Errorf("the queued data length should be 4, but got %d", len(q.Data))
	}
}
