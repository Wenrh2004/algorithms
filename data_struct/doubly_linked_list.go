/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: doubly_linked_list.go
 * Last date: 2023/12/16 下午4:07
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/16 16:7:29.
 */

package data_struct

type DoublyLinkedListNode struct {
	value int
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

// Inserts a node at the end of a doubly linked list
func (list *DoublyLinkedList) AddNode(value int) {
	newNode := &DoublyLinkedListNode{
		value: value,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		current := list.tail
		current.next = newNode
		newNode.prev = current
		list.tail = newNode
	}
}
