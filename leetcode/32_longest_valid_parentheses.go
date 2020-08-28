package leetcode

import (
	"errors"
	"fmt"
)

//32. Longest Valid Parentheses
//Hard
//
//3712
//
//139
//
//Add to List
//
//Share
//Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
//Example 1:
//
//Input: "(()"
//Output: 2
//Explanation: The longest valid parentheses substring is "()"
//Example 2:
//
//Input: ")()())"
//Output: 4
//Explanation: The longest valid parentheses substring is "()()"
//
// 这道题目看着才是正常的题目，之前的很多题目感觉要写很多的判断和小代码，没有那种一个方案就能解决
// 的方式，例如递归了，动态规划啥的
// ( 40 ) 41

func LongestValidParentheses() {
	source := "()(()"
	result := doLongestValidParentheses(source)
	fmt.Printf("结果是:%d\n", result)
}

func doLongestValidParentheses(source string) int {
	stack := ParenthesesStack{
		elements:     make([]int, 0),
		currentIndex: 0,
	}

}

type ParenthesesStack struct {
	elements     []int
	currentIndex int
}

func (p *ParenthesesStack) push(element int) {
	p.elements = append(p.elements, element)
	p.currentIndex = p.currentIndex + 1
}

func (p *ParenthesesStack) pop() (int, error) {
	if p.currentIndex > 0 {
		p.currentIndex = p.currentIndex - 1
		return p.elements[p.currentIndex], nil
	}
	return -1, errors.New("越界了")
}
