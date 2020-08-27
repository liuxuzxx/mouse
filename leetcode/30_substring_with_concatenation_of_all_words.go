package leetcode

import "fmt"

//30. Substring with Concatenation of All Words
//Hard
//
//918
//
//1282
//
//Add to List
//
//Share
//You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
//
//
//Example 1:
//
//Input:
//s = "barfoothefoobarman",
//words = ["foo","bar"]
//Output: [0,9]
//Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
//The output order does not matter, returning [9,0] is fine too.
//Example 2:
//
//Input:
//s = "wordgoodgoodgoodbestword",
//words = ["word","good","best","word"]
//Output: []
//
// 我首先想到的一个方案就是：对words进行组合，然后转换成了字符串的搜索，
// 但是需要的时间就是O(m!)*O(n)
// 现在采用这种方案：因为长度一样子，所以，直接搜索第一个可见的

func SubstringConcatenationAllWords() {
	source := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	result := doSubstringConcatenationAllWords(source, words)
	for _, value := range result {
		fmt.Printf("结果是:%d\n", value)
	}
}

func doSubstringConcatenationAllWords(s string, words []string) []int {
	wordsLength := len(words)
	if wordsLength <= 0 {
		return []int{}
	}
	singleLength := len(words[0])
	totalLength := wordsLength * singleLength

	wordMap := make(map[string]int, 0)
	for _, value := range words {
		wordMap[value] = wordMap[value] + 1
	}

	sourceRune := []rune(s)

	result := make([]int, 0)
	for index := 0; index <= len(s)-totalLength; index = index + 1 {
		tempRune := sourceRune[index : index+totalLength]

		flag := true
		for i := 0; i < totalLength; i = i + singleLength {
			compareRune := tempRune[i : i+singleLength]
			compareStr := string(compareRune)
			if value, ok := wordMap[compareStr]; ok && value > 0 {
				wordMap[compareStr] = value - 1
			} else {
				flag = false
				break
			}
		}
		wordMap = make(map[string]int, 0)
		for _, value := range words {
			wordMap[value] = wordMap[value] + 1
		}
		if flag {
			result = append(result, index)
		}
	}

	return result
}
