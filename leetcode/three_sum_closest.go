package leetcode

import (
	"fmt"
	"math"
)

//16. 3Sum Closest
//Medium
//Given an array nums of n integers and an integer target,
// find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//Example:
//
//Given array nums = [-1, 2, 1, -4], and target = 1.
//
//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

func ThreeSumClosest() {
	sourceArr := []int{-1, 2, 1, -4}
	if len(sourceArr) < 3 {
		fmt.Println("请传入的数组的数字个数至少是3个")
		return
	}
	quickSort(sourceArr, 0, len(sourceArr)-1)
	closestTuple := doThreeSumClosest(sourceArr, 1)
	fmt.Printf("和最接近:%d \n", closestTuple.first+closestTuple.second+closestTuple.third)

	left, right := binarySearch(sourceArr, -3)
	fmt.Printf("left:right=%d-%d", left, right)
}

func quickSort(sourceArr []int, left int, right int) {
	if left < right {
		splitIndex := partition(sourceArr, right, left)
		quickSort(sourceArr, left, splitIndex-1)
		quickSort(sourceArr, splitIndex+1, right)
	}
}

func partition(sourceArr []int, right int, left int) int {
	benchValue := sourceArr[right]
	splitIndex := left
	for index := left; index < right; index++ {
		if benchValue > sourceArr[index] {
			sourceArr[splitIndex], sourceArr[index] = sourceArr[index], sourceArr[splitIndex]
			splitIndex += 1
		}
	}
	sourceArr[splitIndex], sourceArr[right] = sourceArr[right], sourceArr[splitIndex]
	return splitIndex
}

func doThreeSumClosest(sourceArr []int, target int) ThreeTuple {
	closestValue := math.MaxInt32
	var result ThreeTuple
	for index := 0; index < len(sourceArr)-2; index++ {
		second, third := doTwoSumClosest(sourceArr[index+1:], target-sourceArr[index])
		if abs(target-(sourceArr[index]+second+third)) < closestValue {
			closestValue = abs(target - (sourceArr[index] + second + third))
			result = ThreeTuple{
				first:  sourceArr[index],
				second: second,
				third:  third,
			}
		}
	}
	return result
}

func doTwoSumClosest(sourceArr []int, target int) (int, int) {
	left, right := sourceArr[0], sourceArr[1]
	twoClosestValue := math.MaxInt32
	for leftIndex := 0; leftIndex < len(sourceArr)-1; leftIndex++ {
		for rightIndex := leftIndex + 1; rightIndex < len(sourceArr); rightIndex++ {
			difference := sourceArr[leftIndex] + sourceArr[rightIndex] - target
			if abs(difference) < twoClosestValue {
				twoClosestValue = abs(difference)
				left, right = sourceArr[leftIndex], sourceArr[rightIndex]
			}
		}
	}
	return left, right
}

func abs(source int) int {
	if source < 0 {
		return -1 * source
	}
	return source
}

//
//当二分法查找不到，但是又需要查找一个范围的时候，那肯定就是
//返回两个索引了，一左一右
//
func binarySearch(sourceArr []int, target int) (int, int) {
	if target < sourceArr[0] {
		return 0, 0
	}
	if target > sourceArr[len(sourceArr)-1] {
		return len(sourceArr) - 1, len(sourceArr) - 1
	}
	return doBinarySearch(sourceArr, 0, len(sourceArr)-1, target)
}

func doBinarySearch(sourceArr []int, left int, right int, target int) (int, int) {
	if right-left < 2 {
		return left, right
	}
	middle := (left + right) / 2
	if sourceArr[middle] == target {
		return middle, middle
	} else if sourceArr[middle] > target {
		return doBinarySearch(sourceArr, left, middle, target)
	} else {
		return doBinarySearch(sourceArr, middle, right, target)
	}

}
