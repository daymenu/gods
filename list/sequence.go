package list

import (
	"errors"
)

// 链表顺序存储
var (
	IndexError    = errors.New("index out of range")
	PositionError = errors.New("Incorrect position")
)

// 数据定义
type Student struct {
	No   string
	Name string
	Sex  string
}

// 用切片做链表的顺序存储
var students []Student

// 判断是否为空
func ListEmpty() bool {
	if len(students) == 0 {
		return true
	} else {
		return false
	}
}

// 清空链表
func ClearList() {
	students = students[:0]
}

// 链表长度
func ListLength() int {
	return len(students)
}

// 获得元素
func GetElem(i int) (stu Student, err error) {
	l := len(students)
	if i < 1 || i > l {
		return stu, IndexError
	}
	return students[i-1], nil
}

// 插入元素
func ListInsert(i int, stu Student) error {
	l := len(students)
	if i < 1 || i > l+1 {
		return PositionError
	}
	students = append(students, stu)
	if i == l+1 {
		return nil
	}

	for j := len(students); j > i; j-- {
		students[j+1] = students[j]
	}
	students[i] = stu
	return nil
}

// 删除元素
func ListDelete(i int) (stu Student, err error) {
	l := len(students)
	if i < 1 || i > l {
		return stu, PositionError
	}
	stu = students[i-1]
	for j := i; j < l; j++ {
		students[j-1] = students[j]
	}
	students = students[:l-1]
	return
}

// 搜索元素
func LocateElem(no string) (index int, ok bool) {
	for i, s := range students {
		if s.No == no {
			return i+1, true
		}
	}
	return index, ok
}
