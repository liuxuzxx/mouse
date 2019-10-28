package leetcode

import (
	"fmt"
	"strconv"
)

//Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
//Example 1:
//
//Input: 121
//Output: true
//Example 2:
//
//Input: -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
//Example 3:
//
//Input: 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//Follow up:
//
//Could you solve it without converting the integer to a string?
//有一种方案：就是利用上面有个题目，进行倒转操作，得到的数字和原来的数字比较，相等，说明对称，不相等说明不对称
func PalindromeNumber() {
	source := 10
	result := doOperation(source)
	if result {
		fmt.Println(source, " is a palindrome number!")
	} else {
		fmt.Println(source, " no!")
	}
}

func doOperation(source int) bool {
	if source < 0 {
		return false
	}
	sourceStr := strconv.Itoa(source)
	length := len(sourceStr)
	middle := length / 2
	for index := 0; index < middle; index++ {
		if sourceStr[index] != sourceStr[length-1-index] {
			return false
		}
	}
	return true
}
