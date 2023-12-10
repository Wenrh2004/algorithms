/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: binarySearch_test.go
 * Last date: 2023/12/11 上午12:24
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/11 0:24:32.
 */

package Search

import "testing"

func TestBinarySearch(t *testing.T) {
	// 测试用例
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	
	// 测试存在的目标值
	target := 11
	expected := 5 // 目标值 11 在数组中的索引为 5
	if result := binarySearch(arr, target); result != expected {
		t.Errorf("期望目标值 %d 的索引为 %d，但实际得到索引 %d", target, expected, result)
	}
	
	// 测试不存在的目标值
	target = 8
	expected = -1 // 目标值 8 未在数组中找到，返回 -1
	if result := binarySearch(arr, target); result != expected {
		t.Errorf("期望目标值 %d 未在数组中找到，但实际返回索引 %d", target, result)
	}
}
