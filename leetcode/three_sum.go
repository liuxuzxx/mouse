package leetcode

import "fmt"

//15. 3Sum
//Medium
//Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
//Note:
//
//The solution set must not contain duplicate triplets.
//
//Example:
//
//Given array nums = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//我个人的思路还是就是按照一个去找，然后，继续查找就行
func ThreeSum() {
	sourceArr := []int{-1, 0, 1, 2, -1, -4}
	result := doThreeSum(sourceArr, 0)
	for _, value := range result {
		fmt.Printf("[%d,%d,%d]\n", value.first, value.second, value.third)
	}
}

func doThreeSum(sourceArr []int, sum int) []ThreeTuple {
	var result = make([]ThreeTuple, 0)
	for index := 0; index < len(sourceArr); index++ {
		twoSum := sum - sourceArr[index]
		twoResult := doTwoSum(sourceArr[index+1:], twoSum, sourceArr[index])
		result = append(result, twoResult...)
	}
	return result
}

func doTwoSum(sourceArr []int, sum int, first int) []ThreeTuple {
	var target = make([]ThreeTuple, 0)
	for index := 0; index < len(sourceArr); index++ {
		right := sum - sourceArr[index]
		for innerIndex := index + 1; innerIndex < len(sourceArr); innerIndex++ {
			if right == sourceArr[innerIndex] {
				target = append(target, ThreeTuple{first: first, second: sourceArr[index], third: right})
			}
		}
	}
	return target
}

type ThreeTuple struct {
	first  int
	second int
	third  int
}
