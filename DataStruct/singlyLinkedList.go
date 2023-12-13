/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: singlyLinkedList.go
 * Last date: 2023/12/13 下午11:31
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/13 23:31:59.
 */

package DataStruct

import "fmt"

type SinglyLinkedListNode struct {
	value int
	next  *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
	}
}

func (receiver SinglyLinkedList) name() {

}

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

func (list *SinglyLinkedList) Delete() {
	list.head = list.head.next
}

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

// Search target in singly linked list
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

func (list SinglyLinkedList) Size() int {
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

func (list *SinglyLinkedList) Print() {
	node := list.head

	for node != nil {
		fmt.Printf("%v -> ", node.value)
		// fmt.Println(node.next)
		node = node.next
	}

	println("nil")
}
