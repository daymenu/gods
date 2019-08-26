package list

// 实现可比较功能

// Comparer 数据接口
type Comparer interface {
	Compare(data interface{}) bool
}
