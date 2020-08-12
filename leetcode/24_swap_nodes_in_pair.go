package leetcode

import "fmt"

//24. Swap Nodes in Pairs
//Medium
//
//2423
//
//174
//
//Add to List
//
//Share
//Given a linked list, swap every two adjacent nodes and return its head.
//
//You may not modify the values in the list's nodes, only nodes itself may be changed.
//
//
//
//Example:
//
//Given 1->2->3->4, you should return the list as 2->1->4->3.
//

func SwapTwoNodes() {
	data := []int{1, 2, 3, 4, 5, 9, 0}
	source := constructListNode(&data)
	result := doSwapTwoNodes(source)
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
}

func doSwapTwoNodes(head *ListNode) *ListNode {
	top := &ListNode{
		Val:  0,
		Next: head,
	}
	prevElement := top
	cursor := prevElement.Next
	for cursor != nil && cursor.Next != nil {
		temp := cursor.Next
		cursor.Next = temp.Next
		temp.Next = cursor
		prevElement.Next = temp
		prevElement = cursor
		cursor = cursor.Next
	}
	return top.Next
}
