/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: singly_linked_list.go
 * Last date: 2023/12/18 上午12:15
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/18 0:15:57.
 */

package data_struct

import "fmt"

// SinglyLinkedListNode ==> The structure of the node in singly linked list.
type SinglyLinkedListNode struct {
	value int
	next  *SinglyLinkedListNode
}

// SinglyLinkedList ==> The structure of the singly linked list.
type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

// NewSinglyLinkedList ==> Create a new singly linked list.
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
	}
}

// Add ==> To add node to the head of the singly linked list.
func (list *SinglyLinkedList) Add(value int) {
	newNode := &SinglyLinkedListNode{
		value: value,
	}

	newNode.next = list.head
	list.head = newNode
}

// Insert ==> To insert node to the earliest node in singly linked list.
func (list *SinglyLinkedList) Insert(value int) {
	newNode := &SinglyLinkedListNode{
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

// Delete ==> To delete the lastest node in singly linked list.
func (list *SinglyLinkedList) Delete() {
	list.head = list.head.next
}

// Remove ==> To remove target node in singly linked list.
func (list *SinglyLinkedList) Remove(target int) {
	node := list.head
	if node == nil {
		return
	}

	if node.value == target {
		list.head = node.next
		return
	}

	prev := list.head
	for prev.next != nil {
		if prev.next.value == target {
			prev.next = prev.next.next
			return
		}
		prev = prev.next
	}
}

// Search ==> To search target node in singly linked list.
func (list *SinglyLinkedList) Search(target int) int {
	node := list.head

	count := 0
	if node == nil {
		return -1
	}

	for node != nil {
		if node.value == target {
			return count
		}
		node = node.next
		count++
	}

	if node == nil {
		return -1
	}

	return count
}

// Size ==> To get the singly linked list's size.
func (list *SinglyLinkedList) Size() int {
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

// Print ==> To print the node in singly linked list
func (list *SinglyLinkedList) Print() {
	node := list.head

	for node != nil {
		fmt.Printf("%v -> ", node.value)
		node = node.next
	}

	println("nil")
}
