package binarysearch

import (
	"strings"
	"testing"
)

type Order struct {
	ID       string
	GoodName string
}

func (o *Order) Compare(data interface{}) int {
	d := data.(*Order)
	return strings.Compare(o.ID, d.ID)
}
func TestTree(t *testing.T) {
	tree := New()
	o := &Order{ID: "J0001", GoodName: "可乐"}
	o1 := &Order{ID: "J0002", GoodName: "雪碧"}
	o2 := &Order{ID: "J0003", GoodName: "茶"}
	o3 := &Order{ID: "J0004", GoodName: "内裤"}
	o4 := &Order{ID: "J0004", GoodName: "内裤1"}
	tree.Insert(o1)
	tree.Insert(o)
	tree.Insert(o2)
	tree.Insert(o3)
	tree.Insert(o4)
	tree.PreOrder()
	// tree.InOrder()
	// tree.PostOrder()
	datas := tree.Find(o4)
	for i := range datas {
		t.Log(datas[i])
	}
	t.Error()
}
