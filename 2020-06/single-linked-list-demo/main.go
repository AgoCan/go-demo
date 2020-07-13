package main

import (
	"errors"
	"fmt"
)

// Node 节点
type Node struct {
	Data interface{}
	Next *Node
}

// ListNode 头节点
type ListNode struct {
	headNode *Node
}

// Length 获取长度
func (list *ListNode) Length() int {
	current := list.headNode
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// IsEmpty 判断是否为空
func (list *ListNode) IsEmpty() bool {
	return list.headNode == nil
}

// GetHeadNode 获取头节点
func (list *ListNode) GetHeadNode() *Node {
	return list.headNode
}

// Add 头部添加元素
func (list *ListNode) Add(data interface{}) {
	node := &Node{Data: data}
	node.Next = list.headNode
	list.headNode = node
}

// Append 末尾添加元素
func (list *ListNode) Append(data interface{}) {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.headNode = node
	} else {
		current := list.headNode
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// Insert 中间插入元素
func (list *ListNode) Insert(index int, data interface{}) {
	node := &Node{Data: data}
	if index < 0 {
		list.Add(data)
	} else if index > list.Length() {
		list.Append(data)
	} else if list.IsEmpty() {
		list.headNode = node
	} else {
		current := list.headNode
		for i := 0; i < index-2; i++ {
			current = current.Next
		}
		node.Next = current.Next
		current.Next = node
	}
}

// P 获取长度
func (list *ListNode) P() {
	current := list.headNode
	fmt.Println(current)
	for current != nil {

		current = current.Next
		if current != nil {
			fmt.Println(current)
		}
	}
	return
}

// Remove 删除指定元素，删除第一个被找到的元素，后面的元素并不会删除
func (list *ListNode) Remove(data interface{}) {
	current := list.headNode
	if current.Data == data {
		list.headNode = current.Next
	} else {
		for current.Next != nil {
			if current.Next.Data == data {
				current.Next = current.Next.Next
			} else {
				current = current.Next
			}
		}
	}
}

// RemoveAtIndex 根据index删除节点
func (list *ListNode) RemoveAtIndex(index int) (node *Node, error error) {
	current := list.headNode
	if index <= 0 {
		current = current.Next
		return current, nil
	} else if index > list.Length() {
		error = errors.New("out of bound")
		return nil, error
	} else {
		for count := 0; count < index-2 && current.Next != nil; count++ {
			current = current.Next
		}
		defer func() {
			current.Next = current.Next.Next
		}()
		return current.Next, nil
	}
}

// GetNodeAtIndex 根据索引查找节点
func (list *ListNode) GetNodeAtIndex(index int) (node *Node, err error) {
	if list.IsEmpty() {
		return nil, nil
	} else if index < 0 {
		return list.headNode, nil
	} else if index > list.Length() {
		return nil, errors.New("out of bound")
	} else {
		current := list.headNode
		count := 0
		for count < index-1 {
			count++
			current = current.Next
		}
		return current, nil
	}
}

// GetNode 获取节点
func (list *ListNode) GetNode(data interface{}) (node *Node) {
	current := list.headNode
	if current == nil {
		return nil
	} else if current.Data == data {
		return current
	} else {
		for current.Next != nil {
			current = current.Next
			if current.Data == data {
				return current
			}
		}
		return nil
	}
}

// Contain 判断节点是否存在
func (list *ListNode) Contain(data interface{}) bool {
	node := list.GetNode(data)
	return node != nil
}

func main() {
	list := ListNode{}
	list.Add("头部")
	// 添加尾部
	list.Append("尾部")
	// 获取长度
	fmt.Println(list.Length())
	// 获取头节点
	fmt.Println(list.GetHeadNode())
	list.P()
}
