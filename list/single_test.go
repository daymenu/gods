package list

import "testing"

func TestSingleInsert(t *testing.T) {
	l := Single()
	l.Insert(1, "aaa")
	l.Insert(2, "bbb")
	l.Insert(3, "ccc")
	// l1 := Single()
	// l1.Insert(1, "ddd")
	// l.Union(l1)
	l.Delete(0)
	// t.Error(l)
	// l.Reverse()
	// t.Error(l.head.next.Data)
	t.Error(l)
}
