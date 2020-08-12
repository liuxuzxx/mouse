package leetcode

import "fmt"

//23. Merge k Sorted Lists
//Hard
//
//5024
//
//296
//
//Add to List
//
//Share
//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
//Example:
//
//Input:
//[
//1->4->5,
//1->3->4,
//2->6
//]
//Output: 1->1->2->3->4->4->5->6
//
// 这个使用最小堆就可以了
//
func MergeKSortedList() {
	first := []int{1, 4, 5}
	second := []int{1, 3, 4}
	third := []int{2, 6}
	sources := make([]*ListNode, 0)
	sources = append(sources, constructListNode(&first))
	sources = append(sources, constructListNode(&second))
	sources = append(sources, constructListNode(&third))

	result := doMergeKSortedList(sources)
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
}

func doMergeKSortedList(sources []*ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cursor := head
	filterSource := make([]*ListNode, 0)
	for _, value := range sources {
		if value != nil {
			filterSource = append(filterSource, value)
		}
	}

	minHeap := buildMinHeap(filterSource)
	tailLength := len(minHeap)

	for tailLength > 0 {
		topElement := minHeap[0]
		cursor.Next = &ListNode{
			Val:  topElement.Val,
			Next: nil,
		}
		cursor = cursor.Next
		if topElement.Next != nil {
			minHeap[0] = topElement.Next
		} else {
			minHeap[0] = minHeap[tailLength-1]
			minHeap[tailLength-1] = nil
			tailLength = tailLength - 1
		}
		minHeap = adjustMinHeap(minHeap, 0, tailLength)
	}
	return head.Next
}

func buildMinHeap(source []*ListNode) []*ListNode {
	length := len(source)
	middle := length/2 - 1
	for index := middle; index >= 0; index = index - 1 {
		adjustMinHeap(source, index, length)
	}
	return source
}

func adjustMinHeap(source []*ListNode, index, length int) []*ListNode {
	left := index*2 + 1
	right := index*2 + 2
	target := index
	if left < length && source[left].Val < source[index].Val {
		target = left
	} else {
		target = index
	}
	if right < length && source[right].Val < source[target].Val {
		target = right
	}
	if target != index {
		temp := source[index]
		source[index] = source[target]
		source[target] = temp
		return adjustMinHeap(source, target, length)
	}
	return source
}
