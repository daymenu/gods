package list

import "testing"

func TestDInsert(t *testing.T) {
	l := Double()
	l.Insert(1, 1)
	l.Insert(2, 2)
	l.Insert(3, 3)
	l.Insert(4, 4)
	t.Error(l)
}
