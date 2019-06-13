package chain

import (
	"testing"
)
func TestNew(t *testing.T){
	q:=New()
	if q.length != 0{
		t.Error("front should nil")
	}
}

