package skiplist

import (
	"fmt"
	"sort"
	"testing"
)

func TestInsert(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(1, 1)
	sl.Insert(2, 2)
	sl.Insert(3, 3)
	sl.Insert(4, 4)
	sl.Insert(5, 5)
	s, err := sl.Find(3, 3)
	fmt.Printf("%#v %v", s, err)
	t.Error()
}

/**
测试随机函数的概率
*/
func TestRandomLevel(t *testing.T) {
	randoms := make(map[int]int)
	randomTime := 10000000
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
