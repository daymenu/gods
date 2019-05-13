package list

import "fmt"

// 链表顺序存储

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
		return stu, fmt.Errorf("元素未找到")
	}
	return students[i-1], nil
}

// 插入元素
func ListInsert(i int, stu Student) error {
	l := len(students)
	if i < 1 || i > l+1 {
		return fmt.Errorf("插入位置不正确")
	}
	students = append(students, stu)
	if i == l {
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
		return stu, fmt.Errorf("插入位置不正确")
	}
	stu = students[i-1]
	for j := i; j < l; j++ {
		students[j-1] = students[j]
	}
	students = students[:l]
	return
}

// 搜索元素
func LocateElem(stu Student) (index int, ok bool) {
	for i, s := range students {
		if s.No == stu.No {
			return i, true
		}
	}
	return index, ok
}
