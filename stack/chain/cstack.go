package chain

import "errors"

var (
	EmptyStack = errors.New("stack is empty")
)

type Element int
type StackNode struct {
	El   Element
	Next *StackNode
}

type LinkStack struct {
	top   *StackNode
	count int
}

func (l *LinkStack) Push(el Element) (err error) {
	s := StackNode{El: el, Next: l.top}
	l.top = &s
	l.count++
	return
}

func (l *LinkStack) Pull() (s *StackNode, err error) {
	if l.count == 0 {
		return s, EmptyStack
	}
	l.count--
	s = l.top
	l.top = l.top.Next
	return
}
