package sequece

// Stack 栈链表要
type Stack struct {
	el  []interface{}
	top int
}

// Push 入栈
func (s *Stack) Push(el interface{}) {
	s.top++
	s.el = append(s.el, el)
}

// Pop 出栈
func (s *Stack) Pop() (el interface{}, ok bool) {
	if s.top < 1 {
		return
	}
	s.top--
	el = s.el[s.top]
	s.el = s.el[:s.top]
	ok = true
	return
}

// Length 长度
func (s *Stack) Length() int {
	return len(s.el)
}

// Clear 清空栈
func (s *Stack) Clear() error {
	s.el = s.el[:0]
	return nil
}
