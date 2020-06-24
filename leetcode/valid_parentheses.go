package leetcode

import "fmt"

//20. Valid Parentheses
//Easy
//
//4685
//
//210
//
//Add to List
//
//Share
//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.
//
//Example 1:
//
//Input: "()"
//Output: true
//Example 2:
//
//Input: "()[]{}"
//Output: true
//Example 3:
//
//Input: "(]"
//Output: false
//Example 4:
//
//Input: "([)]"
//Output: false
//Example 5:
//
//Input: "{[]}"
//Output: true
//类似于语法当中的分析的括号的配对,这个使用栈是个经典的解决方案
//因为栈有个特点就是：后进先出，本身操作系统的栈就有这个配对的关系
//队列的特点是先进先出，没有配对和匹配以及包含的方式
//操作系统当中调用一个函数，如果进入了一个函数，首先是压栈一些东西，等到执行完
//之后，会弹出，然后执行上一个压栈的，可以理解为{[]}这种匹配方式

func ValidParentheses() {
	parentheses := "([)"
	flag := doParentheses(parentheses)
	fmt.Print(flag)
}

func doParentheses(parentheses string) bool {
	stack := Stack{
		content:      []byte(""),
		currentIndex: 0,
	}
	elements := []byte(parentheses)
	for _, value := range elements {
		stack.compare(value)
	}
	return stack.size() == 0
}

type Stack struct {
	content      []byte
	currentIndex int
}

func (s Stack) size() int {
	return s.currentIndex
}

func (s Stack) push(parentheses byte) {
	_ = append(s.content, parentheses)
	s.currentIndex = s.currentIndex + 1
}

func (s Stack) pop() {
	s.currentIndex = s.currentIndex - 1
}

func (s Stack) compare(parentheses byte) {
	if s.currentIndex == 0 {
		s.push(parentheses)
		return
	}
	minNumber := s.content[s.currentIndex-1] - parentheses
	if minNumber > 0 && minNumber < 2 {
		s.currentIndex = s.currentIndex - 1
	} else {
		s.push(parentheses)
	}
}

