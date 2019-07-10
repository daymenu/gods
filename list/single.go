package list

import "fmt"

// SingleElement 列表结点
type SingleElement struct {
	Data interface{}
	next *SingleElement
}

// SingleLink 列表
type SingleLink struct {
	head   *SingleElement
	lenght int
}

// New 新建一个空列表
func New() *SingleLink {
	return &SingleLink{
		head: &SingleElement{},
	}
}

// Length 获取列表的长度
func (l *SingleLink) Length() int {
	return l.lenght
}

// Insert 在列表的第几个位置插入元素
func (l *SingleLink) Insert(i int, data interface{}) error {
	maxLen := l.Length() + 1
	if i < 0 || i > maxLen {
		return ErrIndex
	}
	p := l.head
	for j := 1; j <= maxLen; j++ {
		if i == j {
			p.next = &SingleElement{
				Data: data,
				next: p.next,
			}
			l.lenght++
			p = p.next
			break
		}
		p = p.next
	}
	return nil
}

// Delete 删除制定位置的元素
func (l *SingleLink) Delete(i int) (data interface{}, err error) {
	if i < 0 || i > l.Length() {
		return nil, ErrIndex
	}
	p := l.head
	for j := 1; j <= l.Length(); j++ {
		if i == j {
			data = p.next.Data
			p.next = p.next.next
			break
		}
		p = p.next
	}
	return data, nil
}

// String 实现Stringer接口
func (l *SingleLink) String() string {
	lstr := "\nlist:\n"
	for p := l.head.next; p != nil; {
		lstr += fmt.Sprintln("\t", p.Data)
		p = p.next
	}
	return lstr
}
