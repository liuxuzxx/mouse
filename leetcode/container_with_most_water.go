package leetcode

import "fmt"

//11. Container With Most Water
//Medium
//Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.
//
//The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//
//
//
//Example:
//
//Input: [1,8,6,2,5,4,8,3,7]
//Output: 49

//时间复杂度计算：T(n) = T(n-1) + n o(n2)  n的平方的级别

func ContainerWithMostWater() {
	source := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	left, right, maxValue := doContainerWithMostWater(source, 0)
	fmt.Printf("The result is left:%d right:%d maxValue:%d\n", left, right, maxValue)
}

func doContainerWithMostWater(source []int, startIndex int) (left, right, maxValue int) {
	if startIndex == len(source)-1 {
		return len(source), len(source), 0
	}
	left, right, maxValue = startIndex, 0, 0
	for index := startIndex + 1; index < len(source); index++ {
		tempMax := (index - startIndex) * min(source[index], source[startIndex])
		if tempMax > maxValue {
			right, maxValue = index, tempMax
		}
	}
	subLeft, subRight, subMaxValue := doContainerWithMostWater(source, startIndex+1)
	if maxValue > subMaxValue {
		return left, right, maxValue
	} else {
		return subLeft, subRight, subMaxValue
	}
}

func max(left, right int) int {
	flag := left > right
	if flag {
		return left
	} else {
		return right
	}
}

func min(left, right int) int {
	flag := left > right
	if flag {
		return right
	} else {
		return left
	}
}
