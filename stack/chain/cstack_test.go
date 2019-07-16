package chain

import "testing"

func TestPush(t *testing.T) {
	stack := New()
	stack.Push(1)
	if stack.top.el != 1 {
		t.Error(stack.top, stack.count)
	}
}

func TestPull(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Pop()
	if stack.count != 0 || stack.top != nil {
		t.Error(stack.count, stack.top)
	}
}
