package list

import (
	"testing"
)

type Age int

func (a Age) compare(data interface{}) bool {
	if a == data {
		return true
	}
	return false
}

func TestSingleInsert(t *testing.T) {
	l := Single()
	err := l.Insert(1, Age(1))
	if l.Length() != 1 || l.head.next.Data != Age(1) {
		t.Error(err)
	}
	err = l.Insert(1, Age(2))
	if l.Length() != 2 || l.head.next.Data != Age(2) {
		t.Error(err)
	}
	err = l.Insert(1, Age(3))
	if l.Length() != 3 || l.head.next.Data != Age(3) {
		t.Error(err)
	}
}

func TestSingleDelete(t *testing.T) {
	l := Single()
	if _, err := l.Delete(1); err != ErrIndex {
		t.Error(err)
	}
	l.Insert(1, Age(1))
	l.Insert(1, Age(2))
	if e, err := l.Delete(2); err != nil || e != Age(1) {
		t.Error(err)
	}
}
