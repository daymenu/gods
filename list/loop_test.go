package list

import "testing"

func TestLoopInsert(t *testing.T) {
	l := Loop()
	l.Insert(1, "aaa")
	l.Insert(2, "bbb")
	l.Insert(3, "ccc")
	// l.Delete(l.Length() - 3 + 1)
	// t.Error(l)
	// l.Reverse()
	// t.Error(l.head.next.Data)
	t.Error(l)
}
