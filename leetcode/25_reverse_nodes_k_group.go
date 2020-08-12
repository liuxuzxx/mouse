package leetcode

import "fmt"

//25. Reverse Nodes in k-Group
//Hard
//
//2393
//
//361
//
//Add to List
//
//Share
//Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//Example:
//
//Given this linked list: 1->2->3->4->5
//
//For k = 2, you should return: 2->1->4->3->5
//
//For k = 3, you should return: 3->2->1->4->5
//
//Note:
//
//Only constant extra memory is allowed.
//You may not alter the values in the list's nodes, only nodes itself may be changed.

func ReverseKNodesGroup() {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	head := constructListNode(&source)
	result := doReverseKNodesGroup(head, 5)
	for result != nil {
		fmt.Printf("%d->", result.Val)
		result = result.Next
	}
}

func doReverseKNodesGroup(head *ListNode, k int) *ListNode {
	top := &ListNode{
		Val:  0,
		Next: head,
	}
	startCursor := top
	endCursor := startCursor

	for {
		for index := 0; index < k; index = index + 1 {
			endCursor = endCursor.Next
			if endCursor == nil {
				return top.Next
			}
		}
		temp := &ListNode{
			Val:  0,
			Next: nil,
		}
		tempEndCursor := endCursor.Next
		for pointer := startCursor.Next; pointer != tempEndCursor; {
			tempPointer := pointer.Next
			pointer.Next = temp.Next
			temp.Next = pointer
			pointer = tempPointer
		}
		tempStartCursor := startCursor.Next
		startCursor.Next.Next = tempEndCursor
		startCursor.Next = temp.Next
		startCursor = tempStartCursor
		endCursor = startCursor
	}

}
