package list

import "fmt"

// LoopElement 列表结点
type LoopElement struct {
	Data interface{}
	next *LoopElement
}

// LoopLink 列表
type LoopLink struct {
	head   *LoopElement
	lenght int
}

// Loop 新建一个空列表
func Loop() *LoopLink {
	head := &LoopElement{}
	head.next = head
	return &LoopLink{
		head: head,
	}
}

// Length 获取列表的长度
func (l *LoopLink) Length() int {
	return l.lenght
}

// Insert 在列表的第几个位置插入元素
func (l *LoopLink) Insert(i int, data interface{}) error {
	maxLen := l.Length() + 1
	if i <= 0 || i > maxLen {
		return ErrIndex
	}
	p := l.head
	for j := 1; j < i; j++ {
		p = p.next
	}
	p.next = &LoopElement{
		Data: data,
		next: p.next,
	}
	l.lenght++
	return nil
}

// Delete 删除制定位置的元素
func (l *LoopLink) Delete(i int) (data interface{}, err error) {
	if i <= 0 || i > l.Length() {
		return nil, ErrIndex
	}
	p := l.head
	for j := 1; j < i; j++ {
		p = p.next
	}
	data = p.next.Data
	p.next = p.next.next
	l.lenght--
	return data, nil
}

// String 实现Stringer接口
func (l *LoopLink) String() string {
	lstr := "\nlist:\n"
	for p := l.head.next; p != nil; {
		lstr += fmt.Sprintln("\t", p.Data)
		p = p.next
	}
	return lstr
}

// GetElem 获取指定位置元素
func (l *LoopLink) GetElem(i int) (e *LoopElement, err error) {
	if i <= 0 || i > l.Length() {
		return nil, ErrIndex
	}
	p := l.head
	for j := 1; j <= i; j++ {
		p = p.next
	}
	e = p
	return
}

// Reverse 反转列表
func (l *LoopLink) Reverse() error {
	if l.Length() == 0 {
		return nil
	}
	p := l.head.next
	pre := p.next
	p.next = nil

	for pre != nil {
		t := pre.next
		pre.next = p
		p, pre = pre, t
	}

	l.head.next = p
	return nil
}

// Union 两个链表合并
func (l *LoopLink) Union(sl *LoopLink) error {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = sl.head.next
	return nil
}
