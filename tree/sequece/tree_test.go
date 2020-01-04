package sequece

import "testing"

func TestTree(t *testing.T) {
	tree := New(100)
	tree.Insert(1, &Order{ID: "123", GoodName: "hh"})
	t.Error(tree)
}
