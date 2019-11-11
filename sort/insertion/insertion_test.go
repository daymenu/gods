package insertion

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	a := []int{3, 2, 1, 5, 5, 6, 9, 8}
	Insertion(a)
	if err := isIncrement(a); err != nil {
		t.Log(err)
		t.Fail()
	}
}

func isIncrement(a []int) error {
	aLen := len(a)
	for i := 1; i < aLen; i++ {
		if a[i-1] > a[i] {
			return fmt.Errorf("a[%d] = %d > a[%d] = %d", i-1, a[i-1], i, a[i])
		}
	}
	return nil
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {

		a := []int{3, 2, 1, 5, 5, 6, 9, 8}
		Insertion(a)
	}
}
func BenchmarkMyInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {

		a := []int{3, 2, 1, 5, 5, 6, 9, 8}
		MyInsertion(a)
	}
}
