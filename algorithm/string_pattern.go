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

func kmpMatch(template, pattern string) {
	stateFunc := preprocess(pattern)
	templateRunes := []rune(template)
	patternRunes := []rune(pattern)
	templateLength := len(templateRunes)
	patternLength := len(patternRunes)
	for index := 0; index < templateLength; index = index + 1 {
		for patternIndex := 0; patternIndex < patternLength; {
			if patternRunes[patternIndex] != templateRunes[patternIndex+index] && patternIndex > 0 {
				stateIndex := stateFunc[patternIndex-1]
				patternIndex = patternIndex - stateIndex
			}
			if patternIndex+1 == patternLength {
				fmt.Printf("找到一个有效的位置:%d\n", index)
			}
		}
	}

}

func KMPStringMatch() {
	pattern := "ababaca"
	template := "ababaababaca"
	kmpMatch(pattern, template)
}
