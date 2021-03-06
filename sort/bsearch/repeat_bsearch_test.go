package bsearch

import "testing"

func TestRepeatLBsearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}

	for i := range a {
		lindex := i
		if a[i] == 4 {
			lindex = 3
		}
		j, err := RepeatLBsearch(a, a[i])
		if j != lindex || err != nil {
			t.Error(i, a[i], lindex, j, err)
		}
	}

	j, err := RepeatLBsearch(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", 0, a)
	}

	m, err := RepeatLBsearch(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}

func TestRepeatRBsearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}

	for i := range a {
		lindex := i
		if a[i] == 4 {
			lindex = 5
		}
		j, err := RepeatRBsearch(a, a[i])
		if j != lindex || err != nil {
			t.Error(i, a[i], lindex, j, err)
		}
	}

	j, err := RepeatRBsearch(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", 0, a)
	}

	m, err := RepeatRBsearch(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}
