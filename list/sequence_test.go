package list

import "testing"

func TestAs(t *testing.T) {
	for i := 10; i > 0; i-- {
		t.Error(i)
	}
}
