package trie

import "testing"

func TestInsert(t *testing.T) {
	trie := &Node{}
	trie.Insert("/inn")
	trie.Insert("/int")
	trie.Insert("/tea")
	trie.Insert("/ten")
	trie.Insert("/to")
}
