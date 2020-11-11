package algorithm

import "fmt"

//
// <Introduce to Algorithm>-For string pattern match
//

//
// 计算模式字符串的转移函数
//
func preprocess(pattern string) []int {
	transactionState := make([]int, len(pattern))
	transactionState[0] = 0
	patternRunes := []rune(pattern)
	for index := 1; index < len(patternRunes); index = index + 1 {
		tempState := transactionState[index-1]
		if patternRunes[index] == patternRunes[tempState] {
			transactionState[index] = tempState + 1
		} else {
			for ; tempState > 0; tempState = transactionState[tempState] {
				if patternRunes[index] == patternRunes[tempState] {
					transactionState[index] = tempState + 1
				}
			}
			if tempState == 0 {
				transactionState[index] = 0
			}
		}
	}
	return transactionState
}

func KMPStringMatch() {
	pattern := "abcabcac"
	result := preprocess(pattern)
	fmt.Printf("%v \n", result)
}
