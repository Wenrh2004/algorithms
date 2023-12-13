/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: linkedList_test.go
 * Last date: 2023/12/13 下午10:43
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/13 22:43:14.
 */

package DataStruct

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Create an empty linked list.
	linkedList := NewLinkedList()

	// Add an element in linked list.
	linkedList.insert(1)

	// Verify whether element 1 is contained in the linked list.
	if linkedList.search(1) != 0 {
		t.Errorf("expected 0, got %v", linkedList.search(1))
	}

	// add multiple elements
	linkedList.insert(2)
	linkedList.insert(3)

	// Verify that the linked list contains elements 2 and 3.
	if linkedList.search(2) != 1 {
		t.Errorf("expected 1, got %v", linkedList.search(1))
	}
	if linkedList.search(3) != 2 {
		t.Errorf("expected 2, got %v", linkedList.search(0))
	}

	// Verify whether there are 3 elements in the linked list
	if linkedList.size() != 3 {
		t.Errorf("expected 3, got %v", linkedList.size())
	}

	// delete element
	linkedList.delete()

	// Verify whether there are 2 elements in the linked list
	if linkedList.size() != 2 {
		t.Errorf("expected 2, got %v", linkedList.size())
	}
}
