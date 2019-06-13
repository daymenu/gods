package chain

import "testing"

func TestPush(t *testing.T) {
	stack := new(LinkStack)
	stack.Push(Element(1))
	if stack.top.El != 1 {
		t.Error(stack.top, stack.count)
	}
}

func TestPull(t *testing.T) {
	stack := new(LinkStack)

	stack.Push(Element(1))
	stack.Push(Element(2))
	stack.Pull()
	stack.Pull()
	if stack.count != 0 || stack.top != nil {
		t.Error(stack.count, stack.top)
	}
}
