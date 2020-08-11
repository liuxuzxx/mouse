package leetcode

import "fmt"

//22. Generate Parentheses
//Medium
//
//5521
//
//278
//
//Add to List
//
//Share
//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//For example, given n = 3, a solution set is:
//
//[
//"((()))",
//"(()())",
//"(())()",
//"()(())",
//"()()()"
//]
//
// 就是正好配对的括号
//

func GenerateParentheses() {
	result := doCombineParentheses(3)
	for _, value := range result {
		fmt.Printf("结果是:%s\n", value)
	}
}

func doCombineParentheses(number int) []string {
	singleParentheses := "()"
	if number == 1 {
		return []string{singleParentheses}
	}
	middleResult := doCombineParentheses(number - 1)
	set := make(map[string]interface{})
	for _, value := range middleResult {
		parenthesesBytes := []byte(value)
		for index := range parenthesesBytes {
			temp := make([]byte, 0)
			temp = append(temp, parenthesesBytes[0:index]...)
			temp = append(temp, []byte(singleParentheses)...)
			temp = append(temp, parenthesesBytes[index:]...)
			set[string(temp)] = nil
		}
	}
	result := make([]string, 0)
	for key := range set {
		result = append(result, key)
	}
	return result
}
