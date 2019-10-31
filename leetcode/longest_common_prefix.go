package leetcode

import "fmt"

//14. Longest Common Prefix
//Easy
//
//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//Input: ["flower","flow","flight"]
//Output: "fl"
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
//Note:
//
//All given inputs are in lowercase letters a-z.

func LongestCommonPrefix() {
	sources := []string{"aca", "cba"}
	//result := doLongestCommonPrefix(sources)
	result := doRecursive(sources)
	fmt.Printf("The common prefix string is:%s\n", result)
}

func doLongestCommonPrefix(sources []string) string {
	if len(sources) == 0 {
		return ""
	}
	result := sources[0]
	for index := 1; index < len(sources); index++ {
		result = twoCommonPrefix(result, sources[index])
	}
	return result
}

func twoCommonPrefix(left, right string) string {
	leftRune := []rune(left)
	rightRune := []rune(right)
	minIndex := min(len(leftRune), len(rightRune))
	result := ""
	for index := 0; index < minIndex; index++ {
		if string(leftRune[index]) == string(rightRune[index]) {
			result += string(leftRune[index])
		} else {
			break
		}
	}
	return result
}

//使用分治算法来处理，这速度还是有点慢啊
func doDivideLongestCommonPrefix(sources []string) string {
	if len(sources) == 0 {
		return ""
	}
	return doRecursive(sources)
}

func doRecursive(sources []string) string {
	if len(sources) == 1 {
		return sources[0]
	}
	middle := len(sources) / 2
	left := doRecursive(sources[0:middle])
	right := doRecursive(sources[middle:])
	return twoCommonPrefix(left, right)
}
