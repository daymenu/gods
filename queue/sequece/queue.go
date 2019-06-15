package sequece

import "fmt"

const maxSize = 100

//SqQueue queue data
type SqQueue struct {
	data  [maxSize]int
	front int
	rear  int
}

//New new queue
func New() *SqQueue {
	return &SqQueue{
		front: 0,
		rear:  0,
	}
}

// Length queue length
func (q *SqQueue) Length() int {
	return (q.rear - q.front + maxSize) % maxSize
}

// In in queue
func (q *SqQueue) In(e int) error {
	if (q.rear+1)%maxSize == q.front {
		return fmt.Errorf("quque is full")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % maxSize
	return nil
}

// Out in queue
func (q *SqQueue) Out() (e int, err error) {
	if q.rear == q.front {
		return e, fmt.Errorf("quque is full")
	}
	e = q.data[q.front]
	q.front = (q.front + 1) % maxSize
	return 0, nil
}
