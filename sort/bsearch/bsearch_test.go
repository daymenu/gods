package bsearch

import "testing"

func TestBsearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		j, err := Bsearch(a, a[i])
		if j != i || err != nil {
			t.Error(err)
		}
	}

	j, err := Bsearch(a, 0)
	if j != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v", 0, a)
	}

	m, err := Bsearch(a, 1000000)
	if m != -1 || err == nil {
		t.Errorf("Bsearch %d not in %v %v", m, a, err)
	}
}
