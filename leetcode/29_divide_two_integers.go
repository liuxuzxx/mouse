package leetcode

import "fmt"

//29. Divide Two Integers
//Medium
//
//1263
//
//5462
//
//Add to List
//
//Share
//Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
//
//Return the quotient after dividing dividend by divisor.
//
//The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
//
//Example 1:
//
//Input: dividend = 10, divisor = 3
//Output: 3
//Explanation: 10/3 = truncate(3.33333..) = 3.
//Example 2:
//
//Input: dividend = 7, divisor = -3
//Output: -2
//Explanation: 7/-3 = truncate(-2.33333..) = -2.
//Note:
//
//Both dividend and divisor will be 32-bit signed integers.
//The divisor will never be 0.
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

//
//这个其实就是在考除法的概念,我们暂时定义余数为正余数，不考虑负余数的情况
// 使用辗转相除法就可以了,或者是从计算机的角度来看到加减乘除的做法的规则实现
// 我的做法不是计算机的真正的做法，是数学的做法
//

func DivideTwoIntegers() {
	dividend := 8
	divisor := 2
	result := doDivideTwoIntegers(dividend, divisor)
	fmt.Printf("%d/%d=%d\n", dividend, divisor, result)
}

func doDivideTwoIntegers(dividend int, divisor int) int {
	absDividend := abs(dividend)
	absDivisor := abs(divisor)
	quotient := 0
	for index := 31; index >= 0; index = index - 1 {
		if (absDividend >> index) >= absDivisor {
			quotient = quotient + 1<<index
		}
	}
	if dividend^divisor < 0 {
		return -quotient
	}
	return quotient
}
