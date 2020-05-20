package trie

import (
	"strings"
	"testing"
)

func TestInsert(t *testing.T) {
	trie := &Node{}
	trie.Insert("/inn")
	trie.Insert("/int")
	trie.Insert("/tea")
	trie.Insert("/ten")
	trie.Insert("/to")
	trie.Insert("/tao")
	trie.Insert("/tttt")
	trie.Insert("/t")
}

func TestIndex(t *testing.T) {
	t.Error("te"[0:1])
	t.Error(strings.Index("it", "te"[0:1]))
}
