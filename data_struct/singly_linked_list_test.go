/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: singly_linked_list_test.go
 * Last date: 2023/12/15 下午5:24
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/15 17:24:12.
 */

package data_struct

import (
	"testing"
)

func TestSinglyLinkedListLinkedList(t *testing.T) {
	// Create an empty linked list.
	singlyLinkedList := NewSinglyLinkedList()

	// Add an element in linked list.
	singlyLinkedList.Insert(1)

	singlyLinkedList.Add(2)

	// Verify whether element 1 is contained in the linked list.
	if singlyLinkedList.Search(1) != 1 {
		t.Errorf("expected 0, got %v", singlyLinkedList.Search(1))
	}

	// add multiple elements
	singlyLinkedList.Insert(2)
	singlyLinkedList.Insert(3)

	// Verify that the linked list contains elements 2 and 3.
	if singlyLinkedList.Search(2) != 0 {
		t.Errorf("expected 0, got %v", singlyLinkedList.Search(2))
	}
	if singlyLinkedList.Search(3) != 3 {
		t.Errorf("expected 2, got %v", singlyLinkedList.Search(3))
	}

	// Verify whether there are 3 elements in the linked list
	if singlyLinkedList.Size() != 4 {
		t.Errorf("expected 4, got %v", singlyLinkedList.Size())
	}

	singlyLinkedList.Remove(3)

	singlyLinkedList.Print()
	if singlyLinkedList.Search(3) != -1 {
		t.Errorf("expected -1, got %v", singlyLinkedList.Search(3))
	}

	// delete element
	singlyLinkedList.Delete()
	singlyLinkedList.Print()

	// Verify whether there are 2 elements in the linked list
	if singlyLinkedList.Size() != 2 {
		t.Errorf("expected 1, got %v", singlyLinkedList.Size())
	}
}
