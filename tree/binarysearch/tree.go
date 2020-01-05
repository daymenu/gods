package binarysearch

import "fmt"

// 实现二叉查找树

// Comparer 可比较的数据
type Comparer interface {
	Compare(data interface{}) int
}

// Node tree node
type Node struct {
	data  Comparer
	left  *Node
	right *Node
}

// Tree tree struct
type Tree struct {
	root   *Node
	length int
}

// New new tree
func New() *Tree {
	return &Tree{
		root:   nil,
		length: 0,
	}
}

// Insert 在树插入一个值
func (t *Tree) Insert(data Comparer) error {
	if t.root == nil {
		t.root = &Node{data: data}
		t.length++
		return nil
	}
	p := t.root
	for p != nil {
		if p.data.Compare(data) < 0 {
			if p.left == nil {
				p.left = &Node{data: data}
				t.length++
				return nil
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &Node{data: data}
				t.length++
				return nil
			}
			p = p.right
		}
	}
	return nil
}

// PreOrder 前序遍历
func (t *Tree) PreOrder() {
	p := t.root
	if p == nil {
		return
	}
	p.preOrder()
}

func (n *Node) preOrder() {
	if n.left != nil {
		n.left.preOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.preOrder()
	}
}

// PostOrder 后序遍历
func (t *Tree) PostOrder() {
	p := t.root
	if p == nil {
		return
	}
	p.postOrder()
}

func (n *Node) postOrder() {
	if n.right != nil {
		n.right.preOrder()
	}
	fmt.Println(n.data)
	if n.left != nil {
		n.left.preOrder()
	}
}

// InOrder 中序遍历
func (t *Tree) InOrder() {
	p := t.root
	if p == nil {
		return
	}
	p.inOrder()
}

func (n *Node) inOrder() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.preOrder()
	}
	if n.right != nil {
		n.right.preOrder()
	}
}
