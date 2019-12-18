package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// MaxLevel 跳表的最大层级
const MaxLevel = 16

// DefaultLevel 默认层数
const DefaultLevel = 1

// DefaultLength 默认长度
const DefaultLength = 0

// SkipNode 跳表中的节点
type SkipNode struct {
	v        interface{}
	level    int
	score    int
	forwords []*SkipNode
}

// SkipList 跳表结构体
type SkipList struct {
	head   *SkipNode
	level  int
	length int
}

// NewSkipNode 新建跳表节点
func NewSkipNode(v interface{}, score int, level int) *SkipNode {
	return &SkipNode{
		v:        v,
		level:    level,
		score:    score,
		forwords: make([]*SkipNode, level),
	}
}

// NewSkipList 建立列表
func NewSkipList() *SkipList {
	return &SkipList{
		head:   NewSkipNode(nil, math.MaxInt32, MaxLevel),
		level:  DefaultLevel,
		length: DefaultLength,
	}
}

// Level 跳表层级
func (s *SkipList) Level() int {
	return s.level
}

// Length 跳表长度
func (s *SkipList) Length() int {
	return s.length
}

func randomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for rand.Intn(MaxLevel) < 8 && level < MaxLevel {
		level++
	}
	return level
}

// Insert 跳表插入
func (s *SkipList) Insert(v interface{}, score int) (err error) {

	if nil == v {
		err = fmt.Errorf("Insert: %v is not nil", v)
		return
	}

	p := s.head
	leftForwords := [MaxLevel]*SkipNode{}
	i := MaxLevel - 1
	for ; i >= 0; i-- {
		for nil != p.forwords[i] {
			if p.forwords[i].v == v {
				err = fmt.Errorf("Insert: %v already exists", v)
				return
			}
			if p.forwords[i].score > score {
				leftForwords[i] = p
				break
			}
		}
		if nil == p.forwords[i] {
			leftForwords[i] = p
		}
	}

	randomLevel := randomLevel()

	node := NewSkipNode(v, score, randomLevel)

	for j := 0; j <= randomLevel-1; j++ {
		next := leftForwords[j].forwords[j]
		leftForwords[j].forwords[j] = node
		node.forwords[j] = next
	}
	fmt.Println(leftForwords)
	if randomLevel > s.level {
		s.level = randomLevel
	}

	s.length++
	return
}
