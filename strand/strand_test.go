package strand

import "testing"

func TestIndex(t *testing.T) {
	if Index("hello world", "y") != -1 {
		t.Error(Index("hello world", "world"))
	}
}

func TestGetNext(t *testing.T) {
	type Data struct {
		str  string
		next []int
	}
	data := []Data{
		{"a", []int{0}},
		{"abc", []int{0, 1, 1}},
		{"abcdex", []int{0, 1, 1, 1, 1, 1}},
		{"abcabx", []int{0, 1, 1, 1, 2, 3}},
		{"ababaaaba", []int{0, 1, 1, 2, 3, 4, 2, 2, 3}},
		{"aaaaaaaab", []int{0, 1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for i, d := range data {
		next := getNext([]rune(d.str))
		for j, n := range d.next {
			if n != next[j] {
				t.Errorf("line %d : %v = %v", i, next, d.next)
			}
		}
	}
}

func TestIndexKMP(t *testing.T) {
	if IndexKMP("hello world", "yl") != -1 {
		t.Error(IndexKMP("hello world", "world"))
	}
}
