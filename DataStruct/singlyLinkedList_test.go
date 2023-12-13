/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: singlyLinkedList_test.go
 * Last date: 2023/12/13 下午11:10
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/13 23:10:34.
 */

package DataStruct

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Create an empty linked list.
	linkedList := NewLinkedList()

	// Add an element in linked list.
	linkedList.Insert(1)

	// Verify whether element 1 is contained in the linked list.
	if linkedList.Search(1) != 0 {
		t.Errorf("expected 0, got %v", linkedList.Search(1))
	}

	// add multiple elements
	linkedList.Insert(2)
	linkedList.Insert(3)

	// Verify that the linked list contains elements 2 and 3.
	if linkedList.Search(2) != 1 {
		t.Errorf("expected 1, got %v", linkedList.Search(1))
	}
	if linkedList.Search(3) != 2 {
		t.Errorf("expected 2, got %v", linkedList.Search(0))
	}

	// Verify whether there are 3 elements in the linked list
	if linkedList.Size() != 3 {
		t.Errorf("expected 3, got %v", linkedList.Size())
	}

	linkedList.Remove(3)

	linkedList.Print()
	if linkedList.Search(3) != -1 {
		t.Errorf("expected -1, got %v", linkedList.Search(3))
	}

	// delete element
	linkedList.Delete()

	// Verify whether there are 2 elements in the linked list
	if linkedList.Size() != 1 {
		t.Errorf("expected 1, got %v", linkedList.Size())
	}
}
