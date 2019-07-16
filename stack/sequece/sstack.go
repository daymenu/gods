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
func (s *Stack) Pop() (ok bool, el interface{}) {
	if s.top < 1 {
		return
	}
	s.top--
	el = s.el[s.top]
	s.el = s.el[:s.top]
	return true, el
}
