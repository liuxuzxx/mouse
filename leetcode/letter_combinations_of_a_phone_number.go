package leetcode

import (
	"fmt"
	"strconv"
)

//17. Letter Combinations of a Phone Number
//Medium
//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
//
//
//Example:
//
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//Note:
//
//Although the above answer is in lexicographical order, your answer could be in any order you want.

func LetterCombinationsOfAPhoneNumber() {
	numbers := "3894369"
	result := doLetterCombinationsOfAPhoneNumber(numbers)
	for _, value := range result {
		fmt.Printf("%s,", value)
	}
}

func doLetterCombinationsOfAPhoneNumber(phoneNumber string) []string {
	phoneLetter := map[int64]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	result := make([]string, 0)
	result = append(result, "")
	for _, value := range phoneNumber {
		number, _ := strconv.ParseInt(string(value), 10, 64)
		letters := phoneLetter[number]
		result = combinationTwo(letters, result)
	}
	return result
}

func combinationTwo(letters string, tempCombination []string) []string {
	combination := make([]string, 0)
	for _, value := range letters {
		for _, combinationValue := range tempCombination {
			combination = append(combination, string(value)+combinationValue)
		}
	}
	return combination
}
