package leetcode

import (
	"fmt"
	"strconv"
)

//38. Count and Say
//Easy
//
//1533
//
//10815
//
//Add to List
//
//Share
//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//
//Given an integer n, generate the nth term of the count-and-say sequence. You can do so recursively, in other words from the previous member read off the digits, counting the number of digits in groups of the same digit.
//
//Note: Each term of the sequence of integers will be represented as a string.
//
//
//
//Example 1:
//
//Input: n = 1
//Output: "1"
//Explanation: This is the base case.
//Example 2:
//
//Input: n = 4
//Output: "1211"
//Explanation: For n = 3 the term was "21" in which we have two groups "2" and "1", "2" can be read as "12" which means frequency = 1 and value = 2, the same way "1" is read as "11", so the answer is the concatenation of "12" and "11" which is "1211".
//
//
//Constraints:
//
//1 <= n <= 30

func CountAndSay() {
	number := 10
	result := doCountAndSay(number)
	fmt.Printf("查看结果:%d %s", number, result)
}

func doCountAndSay(number int) string {
	if number == 1 {
		return "1"
	}
	targets := make([]byte, 0)
	targets = append(targets, byte(1))
	for index := 2; index <= number; index = index + 1 {
		targets = recursivelyCount(targets)
	}
	result := ""
	for index := 0; index < len(targets); index = index + 1 {
		result = result + strconv.Itoa(int(targets[index]))
	}
	return result
}

func recursivelyCount(targets []byte) []byte {
	result := make([]byte, 0)
	target := targets[0]
	count := 0
	for i := 0; i < len(targets); i = i + 1 {
		if targets[i] == target {
			count = count + 1
		} else {
			result = append(result, byte(count), target)
			target = targets[i]
			count = 1
		}
	}
	result = append(result, byte(count), target)
	return result
}
