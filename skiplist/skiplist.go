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

// Comparer 数据可比较接口
type Comparer interface {
	Compare(data interface{}) int
}

// SkipNode 跳表中的节点
type SkipNode struct {
	data     Comparer
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
func NewSkipNode(score int, level int, data Comparer) *SkipNode {
	return &SkipNode{
		data:     data,
		level:    level,
		score:    score,
		forwords: make([]*SkipNode, level),
	}
}

// NewSkipList 建立列表
func NewSkipList() *SkipList {
	return &SkipList{
		head:   NewSkipNode(math.MinInt32, MaxLevel, nil),
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
func (s *SkipList) Insert(score int, data Comparer) (err error) {

	if nil == data {
		err = fmt.Errorf("Insert: %v is not nil", data)
		return
	}

	p := s.head
	leftForwords := [MaxLevel]*SkipNode{}
	i := MaxLevel - 1
	for ; i >= 0; i-- {
		for nil != p.forwords[i] {
			if p.forwords[i].data.Compare(data) == 0 {
				err = fmt.Errorf("Insert: %v already exists", data)
				return
			}
			if p.forwords[i].score > score {
				leftForwords[i] = p
				break
			}
			p = p.forwords[i]
		}

		if nil == p.forwords[i] {
			leftForwords[i] = p
		}
	}

	randomLevel := randomLevel()

	node := NewSkipNode(score, randomLevel, data)

	for j := 0; j <= randomLevel-1; j++ {
		next := leftForwords[j].forwords[j]
		leftForwords[j].forwords[j] = node
		node.forwords[j] = next
	}

	if randomLevel > s.level {
		s.level = randomLevel
	}

	s.length++
	return
}

// Find 寻找某个值得元素
func (s *SkipList) Find(score int, data Comparer) (skipnode *SkipNode, err error) {
	p := s.head
	for i := s.level; i >= 0; i-- {
		for nil != p.forwords[i] {
			if p.forwords[i].data.Compare(data) == 0 {
				return p.forwords[i], nil
			} else if p.forwords[i].score > score {
				break
			}
			p = p.forwords[i]
		}
	}
	return
}

func (sn *SkipNode) String() string {
	return fmt.Sprintf("level:%d score:%d value:%v forwords=%v", sn.level, sn.score, &sn.data, sn.forwords)
}
