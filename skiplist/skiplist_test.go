package skiplist

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

type Order struct {
	OrderID   int
	OrderSn   string
	OrderDate int
}

func (o *Order) String() string {
	return fmt.Sprintf("orderSn: %s", o.OrderSn)
}

func (o *Order) Compare(order interface{}) int {
	orderObj, ok := order.(*Order)
	if !ok {
		return -1
	}
	return strings.Compare(o.OrderSn, orderObj.OrderSn)
}

func TestInsert(t *testing.T) {
	sl := NewSkipList()
	o1 := &Order{OrderID: 1, OrderSn: "JD1000", OrderDate: 1577254912}
	sl.Insert(1, o1)
	s, err := sl.Find(1, o1)
	if err != nil {
		t.Error(err)
	}
	ss, ok := s.data.(*Order)
	if !ok || ss.OrderSn != "JD1000" {
		t.Error("find order is error")
	}
}

/**
测试随机函数的概率
*/
func TestRandomLevel(t *testing.T) {
	randoms := make(map[int]int)
	randomTime := 100000
	for i := 0; i < randomTime; i++ {
		randoms[randomLevel()]++
	}
	randKeys := getMapKeys(randoms)
	for _, index := range randKeys {
		t.Logf("%d : %f", index, float64(randoms[index])/float64(randomTime))
	}
	t.Error()
}

func getMapKeys(m map[int]int) []int {
	keys := make([]int, len(m))
	j := 0
	for key := range m {
		keys[j] = key
		j++
	}
	sort.Ints(keys)
	return keys
}
