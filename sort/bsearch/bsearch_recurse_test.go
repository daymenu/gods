package bsearch

import (
	"testing"
)

func TestBsearchResurse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		j, err := Recurse(a, a[i])
		if j != i || err != nil {
			t.Error(err)
		}
	}

	j, err := Recurse(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", j, a)
	}

	m, err := Recurse(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}
