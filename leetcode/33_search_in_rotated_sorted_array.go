package leetcode

import "fmt"

//33. Search in Rotated Sorted Array
//Medium
//
//5660
//
//491
//
//Add to List
//
//Share
//You are given an integer array nums sorted in ascending order, and an integer target.
//
//Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
//If target is found in the array return its index, otherwise, return -1.
//
//
//
//Example 1:
//
//Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
//Example 2:
//
//Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1
//Example 3:
//
//Input: nums = [1], target = 0
//Output: -1
//
//
//Constraints:
//
//1 <= nums.length <= 5000
//-10^4 <= nums[i] <= 10^4
//All values of nums are unique.
//nums is guranteed to be rotated at some pivot.
//-10^4 <= target <= 10^4

func SearchInRotatedSortedArray() {
	numbers := []int{80, 200, 230, 420, 561, 0, 3, 7, 11, 13, 15}
	target := 0
	result := doSearchInRotatedSortedArray(&numbers, target)
	fmt.Printf("查看搜搜结果:%d\n", result)
}

func doSearchInRotatedSortedArray(numberPointer *[]int, target int) int {
	numbers := *numberPointer
	if len(numbers) <= 0 {
		return -1
	}
	firstNumber := numbers[0]
	left, right := 0, len(numbers)-1
	for left <= right {
		middle := (left + right) / 2
		middleNumber := numbers[middle]

		if middleNumber == target {
			return middle
		}

		if firstNumber > target {
			if firstNumber < middleNumber {
				left = middle + 1
			} else if firstNumber > middleNumber && middleNumber > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if firstNumber < target {
			if firstNumber > middleNumber {
				right = middle - 1
			} else if firstNumber < middleNumber && middleNumber > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			return 0
		}
	}
	return -1
}
