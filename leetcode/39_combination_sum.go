package leetcode

import "fmt"

//39. Combination Sum
//Medium
//
//4746
//
//130
//
//Add to List
//
//Share
//Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.
//
//The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.
//
//It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
//
//
//
//Example 1:
//
//Input: candidates = [2,3,6,7], target = 7
//Output: [[2,2,3],[7]]
//Explanation:
//2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
//7 is a candidate, and 7 = 7.
//These are the only two combinations.
//Example 2:
//
//Input: candidates = [2,3,5], target = 8
//Output: [[2,2,2,2],[2,3,3],[3,5]]
//Example 3:
//
//Input: candidates = [2], target = 1
//Output: []
//Example 4:
//
//Input: candidates = [1], target = 1
//Output: [[1]]
//Example 5:
//
//Input: candidates = [1], target = 2
//Output: [[1,1]]
//
//
//Constraints:
//
//1 <= candidates.length <= 30
//1 <= candidates[i] <= 200
//All elements of candidates are distinct.
//1 <= target <= 500

func CombinationSum() {
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	result := doCombinationSum(candidates, target)
	fmt.Printf("查看结果:%d", len(result))
}

func doCombinationSum(candidates []int, target int) [][]int {
	if target <= 0 || len(candidates) == 0 {
		return [][]int{}
	}
	if len(candidates) == 1 && candidates[0] == target {
		return [][]int{{target}}
	}
	first := candidates[0]
	haveFirstResult := doCombinationSum(candidates, target-first)
	noFirstResult := doCombinationSum(candidates[1:], target)

	for index := 0; index < len(haveFirstResult); index = index + 1 {
		haveFirstResult[index] = append(haveFirstResult[index], first)
	}
	sumResult := append(haveFirstResult, noFirstResult...)
	if first == target {
		return append(sumResult, [][]int{{first}}...)
	}
	return sumResult
}
