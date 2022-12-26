package queue

import "errors"

// ArrayQueue 基于切片实现队列
type ArrayQueue struct {
	// 指向队头的下标
	front int64

	// 指向队尾的下标下一个下标
	rear int64

	// 队列的最大容量
	size int64

	// 存放数据的数组容器
	arrs []interface{}
}

// NewArrayQueue 获取指定大小的队列
func NewArrayQueue(size int64) *ArrayQueue {
	return &ArrayQueue{
		size: size + 1,
		arrs: make([]interface{}, size+1),
	}
}

// ISFull 队列是否满
func (q *ArrayQueue) ISFull() bool {
	return (q.rear-q.front+1)%q.size == 0
}

// ISEmpty 队列是否为空
func (q *ArrayQueue) ISEmpty() bool {
	return q.rear == q.front
}

// Add 往队列中加入数据
func (q *ArrayQueue) Add(data interface{}) error {
	if q.ISFull() {
		return errors.New("queue is full")
	}

	q.arrs[q.rear] = data
	q.rear = (q.rear + 1) % q.size

	return nil
}

// Get 从队列中获取数据
func (q *ArrayQueue) Get() (interface{}, error) {
	if q.ISEmpty() {
		return nil, errors.New("queue is empty")
	}

	tmp := q.front
	q.front = (q.front + 1) % q.size

	return q.arrs[tmp], nil
}
