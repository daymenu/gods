package sequece

import "fmt"

// 完全二叉树适合用数组存储

// Order 订单结构体
type Order struct {
	ID       string
	GoodName string
	Date     string
}

// Tree 订单
type Tree struct {
	order []*Order
}

func (t *Tree) String() string {
	orders := ""
	for i := range t.order {
		if t.order[i] == nil {
			continue
		}
		orders += fmt.Sprintf("%+v", *t.order[i])
	}
	return orders
}

// New 新建一颗树
func New(num int) *Tree {
	return &Tree{order: make([]*Order, num)}
}

// Insert 在指定位置插入
func (t *Tree) Insert(i int, order *Order) error {
	t.order[i] = order
	return nil
}
