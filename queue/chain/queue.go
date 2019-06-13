package chain

// Element is data
type Element int

// Queue is Queue Data
type Queue struct {
	e Element
	next *Queue
}
// QueueLink 是队列
type QueueLink struct {
	front,rear *Queue
	length int
}
//New 新建队列
func New() *QueueLink{
	q:= Queue{e:Element(0)}
	return &QueueLink{front:&q,rear:&q,length:0}
}

//Clear 清空队列
func (q *QueueLink) Clear() bool {

	return true
}

//Destroy 销毁队列
func (q *QueueLink) Destroy() bool{
	return true
}

//Empty 队列是否为空
func (q *QueueLink) Empty() bool{
	return true
}

//Head 获取队头元素
func (q *QueueLink) Head() (e Element, err error){
	return
}
//In : 进队列
func (q *QueueLink) In(e Element) (err error){
	node := Queue{e:e}
	q.rear = &node
	q.length++
	return
}

//Out : 出队列
func (q *QueueLink) Out() (e Element, err error){
	return
}
//Length 获取队列长房
func (q *QueueLink) Length() int{
	return 0
}

