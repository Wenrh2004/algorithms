/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: singlyLinkedList_test.go
 * Last date: 2023/12/13 下午11:26
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/13 23:26:44.
 */

package DataStruct

import (
	"testing"
)

func TestSinglyLinkedListLinkedList(t *testing.T) {
	// Create an empty linked list.
	singlyLinkedList := NewSinglyLinkedList()

	// Add an element in linked list.
	singlyLinkedList.Insert(1)

	// Verify whether element 1 is contained in the linked list.
	if singlyLinkedList.Search(1) != 0 {
		t.Errorf("expected 0, got %v", singlyLinkedList.Search(1))
	}

	// add multiple elements
	singlyLinkedList.Insert(2)
	singlyLinkedList.Insert(3)

	// Verify that the linked list contains elements 2 and 3.
	if singlyLinkedList.Search(2) != 1 {
		t.Errorf("expected 1, got %v", singlyLinkedList.Search(1))
	}
	if singlyLinkedList.Search(3) != 2 {
		t.Errorf("expected 2, got %v", singlyLinkedList.Search(0))
	}

	// Verify whether there are 3 elements in the linked list
	if singlyLinkedList.Size() != 3 {
		t.Errorf("expected 3, got %v", singlyLinkedList.Size())
	}

	singlyLinkedList.Remove(3)

	singlyLinkedList.Print()
	if singlyLinkedList.Search(3) != -1 {
		t.Errorf("expected -1, got %v", singlyLinkedList.Search(3))
	}

	// delete element
	singlyLinkedList.Delete()

	// Verify whether there are 2 elements in the linked list
	if singlyLinkedList.Size() != 1 {
		t.Errorf("expected 1, got %v", singlyLinkedList.Size())
	}
}
