package list

import "errors"

var (
	//ErrIndex 超出列表索引
	ErrIndex = errors.New("out of list index")
	//ErrNotFound 没有找到该元素
	ErrNotFound = errors.New("not found this element")
)
