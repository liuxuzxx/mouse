package leetcode

import (
	"fmt"
	"strconv"
)

func PalindromeNumber() {
	source := 1236598
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
