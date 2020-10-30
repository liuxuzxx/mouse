package leetcode

import "fmt"

//40. Combination Sum II
//Medium
//
//2180
//
//75
//
//Add to List
//
//Share
//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.
//
//Each number in candidates may only be used once in the combination.
//
//Note: The solution set must not contain duplicate combinations.
//
//
//
//Example 1:
//
//Input: candidates = [10,1,2,7,6,1,5], target = 8
//Output:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//Example 2:
//
//Input: candidates = [2,5,2,1,2], target = 5
//Output:
//[
//[1,2,2],
//[5]
//]
//
//
//Constraints:
//
//1 <= candidates.length <= 100
//1 <= candidates[i] <= 50
//1 <= target <= 30

func CombinationSumII() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := filterCombination(candidates, target)
	fmt.Printf("%d\n", len(result))
}

func filterCombination(candidates []int, target int) [][]int {
	if target <= 0 || len(candidates) == 0 {
		return [][]int{}
	}
	if len(candidates) == 1 && candidates[0] == target {
		return [][]int{{target}}
	}
	first := candidates[0]
	haveFirstResult := filterCombination(candidates[1:], target-first)
	noFirstResult := filterCombination(candidates[1:], target)

	for index := 0; index < len(haveFirstResult); index = index + 1 {
		haveFirstResult[index] = append(haveFirstResult[index], first)
	}
	sumResult := append(haveFirstResult, noFirstResult...)
	if first == target {
		return append(sumResult, [][]int{{first}}...)
	}
	return sumResult
}
