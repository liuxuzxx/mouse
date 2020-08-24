package leetcode

import "fmt"

//28. Implement strStr()
//Easy
//
//1675
//
//1963
//
//Add to List
//
//Share
//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
//Clarification:
//
//What should we return when needle is an empty string? This is a great question to ask during an interview.
//
//For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
//
//
//
//Constraints:
//
//haystack and needle consist only of lowercase English characters.
//

func StrStr() {
	haystack := ""
	needle := ""
	result := doStrStr(haystack, needle)
	fmt.Printf("查看结果未知:%d\n", result)
}

func doStrStr(haystack string, needle string) int {
	needleRunes := []rune(needle)
	haystackRunes := []rune(haystack)
	needleLength := len(needleRunes)
	haystackLength := len(haystackRunes)
	for index := 0; index <= haystackLength-needleLength; index = index + 1 {
		if compareStr(needleRunes, haystackRunes, index) {
			return index
		} else {
			continue
		}
	}
	return -1
}

func compareStr(source, target []rune, index int) bool {
	for i := 0; i < len(source); i = i + 1 {
		if source[i] != target[index+i] {
			return false
		}
	}
	return true
}
