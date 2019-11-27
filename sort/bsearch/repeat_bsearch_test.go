package bsearch

import "testing"

func TestRepeatLBsearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}

	m, err := RepeatLBsearch(a, 1)
	if m != 0 || err != nil {
		t.Errorf("Bsearch %d not in %v %v", 1, a, err)
	}
	m, err = RepeatLBsearch(a, 2)
	if m != 1 || err != nil {
		t.Errorf("Bsearch %d not in %v %v", 2, a, err)
	}
	m, err = RepeatLBsearch(a, 4)
	if m != 3 || err != nil {
		t.Errorf("Bsearch %d not in %v %v", 4, a, err)
	}
	m, err = RepeatLBsearch(a, 5)
	if m != 6 || err != nil {
		t.Errorf("Bsearch %d not in %v %v", 5, a, err)
	}
	j, err := RepeatLBsearch(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", 0, a)
	}

	m, err = RepeatLBsearch(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}
