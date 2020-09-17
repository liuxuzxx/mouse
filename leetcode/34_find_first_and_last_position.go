package leetcode

import "fmt"

//34. Find First and Last Position of Element in Sorted Array
//Medium
//
//4016
//
//160
//
//Add to List
//
//Share
//Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
//
//Your algorithm's runtime complexity must be in the order of O(log n).
//
//If the target is not found in the array, return [-1, -1].
//
//Example 1:
//
//Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4]
//Example 2:
//
//Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1]
//
//
//Constraints:
//
//0 <= nums.length <= 10^5
//-10^9 <= nums[i] <= 10^9
//nums is a non decreasing array.
//-10^9 <= target <= 10^9

func FindFirstAndLastPosition() {
	nums := []int{1}
	target := 1
	result := doFindFirstAndLastPosition(&nums, target)
	for _, value := range result {
		fmt.Printf("%d ", value)
	}
}

func doFindFirstAndLastPosition(nums *[]int, target int) []int {
	numbers := *nums
	if len(numbers) == 0 {
		return []int{-1, -1}
	}

	first, last := -1, -1

	left, right := 0, len(numbers)-1
	targetIndex := -1
	for left <= right {
		middle := (left + right) / 2
		middleNumber := numbers[middle]
		if middleNumber > target {
			right = middle - 1
		} else if middleNumber < target {
			left = middle + 1
		} else {
			targetIndex = middle
			break
		}
	}

	if targetIndex != -1 {
		for leftIndex := targetIndex; leftIndex >= 0; leftIndex = leftIndex - 1 {
			if numbers[leftIndex] == numbers[targetIndex] {
				first = leftIndex
			} else {
				break
			}
		}
		for rightIndex := targetIndex; rightIndex < len(numbers); rightIndex = rightIndex + 1 {
			if numbers[rightIndex] == numbers[targetIndex] {
				last = rightIndex
			} else {
				break
			}
		}
	}

	return []int{first, last}
}
