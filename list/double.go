package list

import "fmt"

// DoubleElement 列表结点
type DoubleElement struct {
	Data      interface{}
	pre, next *DoubleElement
}

// DoubleLink 列表
type DoubleLink struct {
	head   *DoubleElement
	lenght int
}

// Double 新建双链表
func Double() *DoubleLink {
	return &DoubleLink{
		head: &DoubleElement{},
	}
}

// Length 获取列表的长度
func (d *DoubleLink) Length() int {
	return d.lenght
}

// Insert 在列表的第几个位置插入元素
func (d *DoubleLink) Insert(i int, data interface{}) error {
	maxLen := d.Length() + 1
	if i <= 0 || i > maxLen {
		return ErrIndex
	}
	p := d.head

	for j := 1; j < i; j++ {
		p = p.next
	}
	p.next = &DoubleElement{
		Data: data,
		pre:  p,
		next: p.next,
	}
	d.lenght++
	return nil
}

// Delete 删除制定位置的元素
func (d *DoubleLink) Delete(i int) (data interface{}, err error) {
	if i <= 0 || i > d.Length() {
		return nil, ErrIndex
	}
	p := d.head
	for j := 1; j < i; j++ {
		p = p.next
	}

	if p.next.next != nil {
		p.next.next.pre = p
	}
	p.next = p.next.next
	d.lenght--
	return data, nil
}

// String 实现Stringer接口
func (d *DoubleLink) String() string {
	lstr := "\nlist:\n"
	for p := d.head.next; p != nil; {
		lstr += fmt.Sprintln("\t", p.Data)
		p = p.next
	}
	return lstr
}
