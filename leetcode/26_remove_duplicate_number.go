package leetcode

import "fmt"

//26. Remove Duplicates from Sorted Array
//Easy
//
//2705
//
//5416
//
//Add to List
//
//Share
//Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.
//
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
//Example 1:
//
//Given nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
//
//It doesn't matter what you leave beyond the returned length.
//Example 2:
//
//Given nums = [0,0,1,1,1,2,2,3,3,4],
//
//Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.
//
//It doesn't matter what values are set beyond the returned length.
//Clarification:
//
//Confused why the returned value is an integer but your answer is an array?
//
//Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.
//
//Internally you can think of this:
//
//// nums is passed in by reference. (i.e., without making a copy)
//int len = removeDuplicates(nums);
//
//// any modification to nums in your function would be known by the caller.
//// using the length returned by your function, it prints the first len elements.
//for (int i = 0; i < len; i++) {
//print(nums[i]);
//}
//
// 可以使用 ^ 这个符号进行测试
// 这个是：一样是0 不一样是1
//

func RemoveDuplicateNumber() {
	source := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 5, 5, 5, 9, 9, 9, 9}
	count := doRemoveDuplicateNumber(source)
	fmt.Printf("数量:%d\n", count)
}

func doRemoveDuplicateNumber(source []int) int {
	length := len(source)
	if length <= 0 {
		return 0
	}
	head := &source[0]
	count := 0
	for index := 0; index < length; index = index + 1 {
		if *head^source[index] != 0 {
			count = count + 1
			source[count] = source[index]
			head = &source[index]
		}
	}
	return count + 1
}
