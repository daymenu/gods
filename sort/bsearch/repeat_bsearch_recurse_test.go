package bsearch

import "testing"

func TestRepeatLBsearchRecurseRecurse(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}

	for i := range a {
		lindex := i
		if a[i] == 4 {
			lindex = 3
		}
		j, err := RepeatLBsearchRecurse(a, a[i])
		if j != lindex || err != nil {
			t.Error(i, a[i], lindex, j, err)
		}
	}

	j, err := RepeatLBsearchRecurse(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", 0, a)
	}

	m, err := RepeatLBsearchRecurse(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}

func TestRepeatRBsearchRecurseRecurse(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}

	for i := range a {
		lindex := i
		if a[i] == 4 {
			lindex = 5
		}
		j, err := RepeatRBsearchRecurse(a, a[i])
		if j != lindex || err != nil {
			t.Error(i, a[i], lindex, j, err)
		}
	}

	j, err := RepeatRBsearchRecurse(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", 0, a)
	}

	m, err := RepeatRBsearchRecurse(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}
