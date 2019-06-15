package chain

import (
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	if q.length != 0 {
		t.Error("front should nil")
	}
}

func TestEmpty(t *testing.T) {
	q := New()
	if !q.Empty() {
		t.Error("error should empty")
	}
}

func TestIn(t *testing.T) {
	q := New()
	q.In(Element(1))
	if q.front.next.e != 1 {
		t.Errorf("element is %d", q.front.next.e)
	}
}

func TestOut(t *testing.T) {
	q := New()
	q.In(Element(1))
	if q.front.next.e != 1 {
		t.Errorf("element is %d", q.front.next.e)
	}
	e, err := q.Out()
	if err != nil {
		t.Error(err)
	}
	if e != 1 {
		t.Error("e is not 1")
	}

	if q.Length() != 0 {
		t.Error("queue should empty")
	}
}
