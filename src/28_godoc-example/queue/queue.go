package queue

// Queue  定义队列
type Queue []interface{}

// Push 往队列中添加元素
func (queue *Queue) Push(element ...interface{}) {
	*queue = append(*queue, element...)
}

// Pop 删除队列中的元素
func (queue *Queue) Pop() interface{} {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

