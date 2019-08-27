package sequece

import "testing"

func TestPush(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	if stack.el[0] != 1 {
		t.Fatal("push stack is not success")
	}
}

func TestPull(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	if e, ok := stack.Pop(); !ok || e != 1 {
		t.Fatal("want 1 but get ", e)
	}
}

func TestClear(t *testing.T) {
	// 测试为空
	stack := new(Stack)
	stack.Clear()
	// 测试不为空的清空
	stack.Push(1)
	stack.Push(2)
	stack.Clear()

	if len(stack.el) != 0 {
		t.Fatal("clear stack is not success")
	}
}
