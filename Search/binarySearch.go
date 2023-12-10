/*
 * Copyright (c) 2023.
 * Project: algorithms
 * File: binarySearch.go
 * Last date: 2023/12/11 上午12:28
 * Developer: KingYen
 *
 * Created by KingYen on 2023/12/11 0:28:39.
 */

package Search

/*
 * binarySearch 二分查找
 *
 * input 输入
 * 输入
 * arr 		｜	输入数据的有序集合
 * target	｜	目标值
 *
 * output 输出
 * -1		｜	未找到目标值
 * other	｜	目标元素下标索引
 */
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	
	return -1
}
