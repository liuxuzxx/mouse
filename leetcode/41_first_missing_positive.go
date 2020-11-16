package leetcode

import "fmt"

//41. First Missing Positive
//Hard
//
//4636
//
//894
//
//Add to List
//
//Share
//Given an unsorted integer array nums, find the smallest missing positive integer.
//
//Follow up: Could you implement an algorithm that runs in O(n) time and uses constant extra space.?
//
//
//
//Example 1:
//
//Input: nums = [1,2,0]
//Output: 3
//Example 2:
//
//Input: nums = [3,4,-1,1]
//Output: 2
//Example 3:
//
//Input: nums = [7,8,9,11,12]
//Output: 1
//
//
//Constraints:
//
//0 <= nums.length <= 300
//-231 <= nums[i] <= 231 - 1

func FirstMissingPositive() {
	nums := []int{1, 1}
	result := doFirstMissingPositive(nums)
	fmt.Printf("查看结果:%d\n", result)
}

func doFirstMissingPositive(nums []int) int {
	numsLength := len(nums)
	for index := 0; index < numsLength; {
		temp := nums[index]
		if temp < 1 || temp > numsLength {
			nums[index] = -1
			index = index + 1
			continue
		}
		if temp == index+1 {
			index = index + 1
			continue
		}
		if temp == nums[temp-1] {
			nums[index] = -1
			index = index + 1
			continue
		}
		middleTemp := nums[temp-1]
		nums[temp-1] = temp
		nums[index] = middleTemp
	}

	for index, value := range nums {
		if value < 1 {
			return index + 1
		}
	}
	if numsLength == 0 {
		return 1
	}
	return numsLength + 1
}
