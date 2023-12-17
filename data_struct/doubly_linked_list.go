/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: doubly_linked_list.go
 * Last date: 2023/12/18 上午1:04
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/18 1:4:51.
 */

package data_struct

import "fmt"

// DoublyLinkedListNode ==> The structure of the node in doubly linked list.
type DoublyLinkedListNode struct {
	value int
	next  *DoublyLinkedListNode
	prev  *DoublyLinkedListNode
}

// DoublyLinkedList ==> The structure of a doubly linked list.
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

// NewDoublyLinkedList ==> To creates a new doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

// Add ==> To add a node to the head of doubly linked list.
func (list *DoublyLinkedList) Add(value int) {
	newNode := &DoublyLinkedListNode{
		value: value,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

// Insert ==> To insert a node at the end of a doubly linked list.
func (list *DoublyLinkedList) Insert(value int) {
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

func (list *DoublyLinkedList) Display() {
	node := list.head
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Println()
}
