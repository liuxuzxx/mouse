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
	haystack := "hello"
	needle := "ll"
	result := doStrStr(haystack, needle)
	fmt.Printf("查看结果未知:%d\n", result)
}

func doStrStr(haystack string, needle string) int {
	if len(needle) == 0 || len(haystack) == 0 {
		return -1
	}
	needleByte := []byte(needle)
	haystackByte := []byte(haystack)
	count := -1
	for index, value := range haystackByte {
		if value == needleByte[0] {
			count = index
		}
		for i := 1; i < len(needleByte); i = i + 1 {
			if haystackByte[index+i] != needleByte[i] {
				count = -1
			}
		}
		if count != -1 {
			break
		}
	}
	return count
}
