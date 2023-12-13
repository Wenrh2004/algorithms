/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: linkedList.go
 * Last date: 2023/12/13 下午10:48
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/13 22:48:34.
 */

package DataStruct

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
	}
}

func (list *LinkedList) insert(value int) {
	newNode := &Node{
		value: value,
	}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) delete() {
	list.head = list.head.next
}

// Search target in singly linked list
func (list *LinkedList) search(target int) int {
	node := list.head

	count := 0
	for node == nil {
		return 0
	}

	for node != nil {
		if node.value == target {
			return count
		}
		node = node.next
		count++
	}

	return count
}

func (list LinkedList) size() int {
	node := list.head
	if node == nil {
		panic("nil")
	}

	count := 0
	for node != nil {
		count++
		node = node.next
	}

	return count
}

// func (l *LinkedList) Print() {
// 	node := l.head
// 	for node != nil {
// 		fmt.Printf("%v ", node.value)
// 		fmt.Println(node.next)
// 		node = node.next
// 	}
// }
