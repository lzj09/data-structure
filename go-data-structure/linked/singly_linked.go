package linked

import "errors"

// SinglyLinkedList 单向链表
type SinglyLinkedList struct {
	// head 链表的表头，不存在数据
	head *node

	// num 记录节点个数
	num int64
}

// NewSinglyLinkedList 获取单向链表
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: &node{},
		num:  0,
	}
}

// Add 往链表中加入数据
func (sl *SinglyLinkedList) Add(data interface{}) int64 {
	tmp := sl.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	tmp.next = &node{
		data: data,
	}
	res := sl.num
	sl.num++

	return res
}

// IndexOf 获取数据的索引，未找到则返回-1
func (sl *SinglyLinkedList) IndexOf(data interface{}) int64 {
	tmp := sl.head
	for i := 0; tmp.next != nil; i++ {
		tmp = tmp.next
		if tmp.data == data {
			return int64(i)
		}
	}

	return -1
}

// Get 获取索引获取数据
func (sl *SinglyLinkedList) Get(i int64) (interface{}, error) {
	n := i + 1
	if n > sl.num || i < 0 {
		return nil, errors.New("超出索引范围")
	}

	tmp := sl.head
	var j int64 = 0
	for ; j < n; j++ {
		tmp = tmp.next
	}

	return tmp.data, nil
}

// Remove 根据索引位置删除数据
func (sl *SinglyLinkedList) Remove(i int64) (interface{}, error) {
	n := i + 1
	if n > sl.num || i < 0 {
		return nil, errors.New("超出索引范围")
	}

	// 先找到该索引前一位节点
	tmp := sl.head
	var j int64 = 0
	for ; j < i; j++ {
		tmp = tmp.next
	}
	delNode := tmp.next
	tmp.next = delNode.next
	sl.num--

	return delNode.data, nil
}

// Update 根据索引更新节点数据
func (sl *SinglyLinkedList) Update(i int64, data interface{}) error {
	n := i + 1
	if n > sl.num || i < 0 {
		return errors.New("超出索引范围")
	}

	// 先找到该索引前一位节点
	tmp := sl.head
	var j int64 = 0
	for ; j < i; j++ {
		tmp = tmp.next
	}

	// 需要更新的节点
	updateNode := tmp.next
	updateNode.data = data

	return nil
}

// Clear 清空链表
func (sl *SinglyLinkedList) Clear() {
	sl.head.next = nil
	sl.num = 0
}

// ISEmpty 判断链表是否为空
func (sl *SinglyLinkedList) ISEmpty() bool {
	return sl.head.next == nil
}

// Size 链表节点数
func (sl *SinglyLinkedList) Size() int64 {
	return sl.num
}

// node 单向链表中的数据节点
type node struct {
	// data 节点中存储的数据
	data interface{}

	// next 指向下一个节点的地址
	next *node
}
