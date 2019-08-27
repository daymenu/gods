package sequece

import "errors"

const (
	// MaxSize 最大长度
	MaxSize = 100
)

var (
	//ErrEmpty 队列为错误定义
	ErrEmpty = errors.New("queue is empty")
	//ErrFull 队列为错误定义
	ErrFull = errors.New("queue is full")
)

// SqQuece 队列结构体
type SqQuece struct {
	el    []interface{}
	front int
	rear  int
}

// New 新建队列
func New() *SqQuece {
	return new(SqQuece)
}

// Enqueue 入队
func (q *SqQuece) Enqueue(el interface{}) error {
	if q.rear == MaxSize && q.front != 1 {
		// 数据移动
		moveNum := MaxSize - q.front - 1
		for i := 1; i <= moveNum; i++ {
			q.el[i] = q.el[q.front]
		}
	} else {
		return ErrFull
	}
	q.el[q.rear] = el
	q.rear++
	return nil
}

// Dequeue 出队
func (q *SqQuece) Dequeue() (interface{}, error) {
	if q.rear == q.front {
		return nil, ErrEmpty
	}
	el := q.el[q.front]
	q.front++
	return el, nil
}

// Length 长度
func (q *SqQuece) Length() int {
	return q.rear - q.front
}
