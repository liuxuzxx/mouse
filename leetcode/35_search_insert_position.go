package leetcode

import "fmt"

//35. Search Insert Position
//Easy
//
//2676
//
//265
//
//Add to List
//
//Share
//Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//
//Example 1:
//
//Input: [1,3,5,6], 5
//Output: 2
//Example 2:
//
//Input: [1,3,5,6], 2
//Output: 1
//Example 3:
//
//Input: [1,3,5,6], 7
//Output: 4
//Example 4:
//
//Input: [1,3,5,6], 0
//Output: 0

func SearchInsertPosition() {
	nums := []int{1, 3, 5, 6}
	target := 0
	result := doSearchInsertPosition(&nums, target)
	fmt.Printf("结果是:%d\n", result)
}

func doSearchInsertPosition(nums *[]int, target int) int {
	numbers := *nums
	left, right := 0, len(numbers)-1
	for left <= right {
		middle := (left + right) / 2
		middleNumber := numbers[middle]
		if middleNumber < target {
			left = middle + 1
		} else if middleNumber > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return left
}
