package trie

import (
	"fmt"
	"strings"
)

// Node node
type Node struct {
	word      string
	fullWord  string
	indices   string
	childrens []*Node
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func findLongestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

// Insert insert
func (t *Node) Insert(word string) {
	// 插入根节点
	if len(t.word) == 0 && len(t.childrens) == 0 {
		t.word = word
		t.fullWord = word
		return
	}
walk:
	for {
		//word 等于0就结束循环
		if len(word) == 0 {
			return
		}
		// 寻找最长前缀
		i := findLongestCommonPrefix(t.word, word)
		// 将该节点变为公共节点 并将该节点的差异后缀变为子节点
		if i < len(t.word) {
			child := Node{
				word:      t.word[i:],
				childrens: t.childrens,
				indices:   t.indices,
			}

			t.childrens = []*Node{&child}
			t.indices = fmt.Sprintf("%c", t.word[i])
			t.word = t.word[:i]
			t.fullWord = t.fullWord
		}
		// 插入新节点
		if i < len(word) {
			hasCommon := false
			for j := 0; j < len(t.indices); j++ {
				if word[i] == t.indices[j] {
					hasCommon = true
					break
				}
			}
			word = word[i:]
			if hasCommon {
				index := strings.Index(t.indices, fmt.Sprintf("%c", word[0]))
				if index != -1 {
					t = t.childrens[index]
					continue walk
				}
			} else {
				child := Node{
					word: word,
				}
				t.indices = fmt.Sprintf("%s%c", t.indices, word[0])
				t.childrens = append(t.childrens, &child)
			}
		}
		return
	}

}

// Search search
func (t *Node) Search(word string) (node *Node, err error) {
	return
}
