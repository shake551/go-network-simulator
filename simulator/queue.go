package simulator

import "fmt"

type Event struct {
	ID   int
	Type string
	Time int64
}

type EventQueue struct {
	Data    []Event
	MaxSize int
}

func (q *EventQueue) IsEmpty() bool {
	return len(q.Data) == 0
}

func (q *EventQueue) IsFull() bool {
	return len(q.Data) == q.MaxSize
}

func (q *EventQueue) Enqueue(e Event) {
	if q.IsFull() {
		fmt.Println("queue is full")
		return
	}

	q.Data = append(q.Data, e)
}

func (q *EventQueue) Dequeue() Event {
	if q.IsEmpty() {
		fmt.Println("queue is empty")
		return Event{}
	}

	eventQueue := q.Data[0]
	q.Data = q.Data[1:]

	return eventQueue
}
