package main

import (
	"fmt"
	"mouse/book"
	"mouse/leetcode"
	"mouse/redis"
)

/**
 * 还是需要紧跟语言啊，毕竟，语言挺重要，是表达的唯一途径
 * go run main.go 是运行
 * go build main.go 是编译成所在系统的可执行文件
 * go clean 类似于mvn clean，就是清除已经编译好的可执行的文件一类的东西
 */
func main() {
	leetcode.ValidSudoku()
}

func redisFound() {
	redis.PhoneCache()
	redis.PhoneNumberMD5()
}

func goBook() {
	book.IntegerType()
	book.CompareFloat()
	book.CalMoney(13)
	book.ComplexNumbers()
	book.StringFound()
	book.ShowLanguage()
	book.ShowWeek()
	book.PartInitial()
	book.Sha256()
	book.SliceMonth()
	book.CompareSlice()
	book.Reverse()
	book.AppendElement()
	book.InitialMap()
	book.AddressSliceAndMap()
	book.ConvertToJSON()
	err := book.ReadFile("")
	if err != nil {
		fmt.Println(err)
	}
	book.DivideNumber(10)
	book.FunctionPoint()
	name := book.DeferOperation()
	fmt.Printf("茶案名字:%s\n", name)
	book.OperationArea()
	book.RemoteStrategy()
	book.Pipeline()
	book.BufferChannel()
	//book.Goroutine()
	//book.Server()
}

func leetCode() {
	leetcode.PalindromeNumber()
	leetcode.RegularExpressionMatching()
	leetcode.ContainerWithMostWater()
	leetcode.IntegerToRoman()
	leetcode.RomanToInteger()
	leetcode.LongestCommonPrefix()

	//article.SpiderIdioms()
	leetcode.ThreeSum()
	leetcode.ThreeSumClosest()
	leetcode.LetterCombinationsOfAPhoneNumber()
	leetcode.FourSum()
	leetcode.RemoteNthNodeFromEndOfList(2)
	leetcode.ValidParentheses()
	leetcode.MergeTwoSortedList()
	leetcode.GenerateParentheses()
	leetcode.MergeKSortedList()
	leetcode.SwapTwoNodes()
	leetcode.ReverseKNodesGroup()
	leetcode.RemoveDuplicateNumber()
	leetcode.RemoveElements()
	leetcode.StrStr()
	leetcode.DivideTwoIntegers()
	leetcode.SubstringConcatenationAllWords()
	leetcode.LongestValidParentheses()
	leetcode.SearchInRotatedSortedArray()
	leetcode.FindFirstAndLastPosition()
	leetcode.SearchInsertPosition()
	leetcode.ValidSudoku()
}
