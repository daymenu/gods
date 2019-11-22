package main

import "testing"

func TestBsearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	i, e := Bsearch(a, 6)
	if i != 2 {
		t.Error(i, e)
	}
}
