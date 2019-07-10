package list

import "testing"

func TestInsert(t *testing.T) {
	l := New()
	l.Insert(1, "aaa")
	l.Insert(2, "bbb")
	l.Delete(2)
	// t.Error(l.head.next.Data)
	t.Error(l)
}
