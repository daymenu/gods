package sequece

type Element int

type Stack struct {
	el  []Element
	top int
}

// 入栈
func (s *Stack) Push(el Element) {
	s.top++
	s.el = append(s.el, el)
}

// 出栈
func (s *Stack) Pull() (ok bool, el Element) {
	if s.top < 1 {
		return
	}
	s.top--
	el = s.el[s.top]
	s.el = s.el[:s.top]
	return true, el
}
