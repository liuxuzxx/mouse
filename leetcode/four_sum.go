package leetcode

import "fmt"

//18. 4Sum
//Medium
//Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
//Note:
//
//The solution set must not contain duplicate quadruplets.
//
//Example:
//
//Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A solution set is:
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]

func FourSum() {
	source := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := doFourSum(source, target)
	for _, value := range result {
		fmt.Printf("%d %d %d %d\n", value.three.first, value.three.second, value.three.third, value.forth)
	}
}

func doFourSum(source []int, target int) []FourTuple {
	result := make([]FourTuple, 0)
	for index, value := range source {
		tempTarget := target - value
		threeTuples := doThreeSum(source[index+1:], tempTarget)
		if len(threeTuples) > 0 {
			for _, threeElement := range threeTuples {
				result = append(result, FourTuple{
					three: threeElement,
					forth: value,
				})
			}
		}
	}
	return result
}

type FourTuple struct {
	three ThreeTuple
	forth int
}
