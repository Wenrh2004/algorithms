/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: douly_linked_list_test.go
 * Last date: 2023/12/18 下午8:43
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/18 20:43:54.
 */

package data_struct

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	doublyLinkedList := NewDoublyLinkedList()

	doublyLinkedList.Add(1)
	doublyLinkedList.Add(2)
	doublyLinkedList.Insert(1)
	doublyLinkedList.Insert(2)
	doublyLinkedList.Insert(3)
	println(doublyLinkedList.Search(3))
	doublyLinkedList.Remove()
	doublyLinkedList.Delete()
	doublyLinkedList.Display()
}
