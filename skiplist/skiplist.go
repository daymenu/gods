package skiplist

// MaxLevel 跳表的最大层级
const MaxLevel = 16

// SkipNode 跳表中的节点
type SkipNode struct {
	v     interface{}
	level int
	scroe int
	right *SkipNode
	down  *SkipNode
}

// SkipList 跳表结构体
type SkipList struct {
	head   *SkipNode
	level  int
	length int
}

// NewSkipNode 新建跳表节点
func NewSkipNode(v interface{}, scroe, level int) *SkipNode {
	return &SkipNode{v: v, level: level, scroe: scroe, right: nil, down: nil}
}

// NewSkipList 建立列表
func NewSkipList() *SkipList {
	return &SkipList{nil, 1, 0}
}

// Level 跳表层级
func (s *SkipList) Level() int {
	return s.level
}

// Length 跳表长度
func (s *SkipList) Length() int {
	return s.length
}
