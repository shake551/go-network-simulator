package simulator

import "fmt"

type Event struct {
	Type          string
	Time          float64
	InServiceTime float64 // optional: when event type is finish, this field needed
}

type EventQueue struct {
	Data    []Event
	MaxSize int
}

func (q *EventQueue) IsEmpty() bool {
	return len(q.Data) == 0
}

func (q *EventQueue) IsFull() bool {
	return len(q.Data) >= q.MaxSize
}

func (q *EventQueue) Enqueue(e Event) error {
	if q.IsFull() {
		fmt.Println("queue is full")
		return fmt.Errorf("queue is full")
	}

	q.Data = append(q.Data, e)
	return nil
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
