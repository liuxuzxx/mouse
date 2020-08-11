package leetcode

import "fmt"

//21. Merge Two Sorted Lists
//Easy
//
//4570
//
//610
//
//Add to List
//
//Share
//Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

func MergeTwoSortedList() {
	left := []int{1, 2, 4}
	right := []int{1, 3, 4, 5}
	leftList := constructListNode(&left)
	rightList := constructListNode(&right)
	result := doMerge(leftList, rightList)
	for result != nil {
		fmt.Printf("%d,", result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func doMerge(left, right *ListNode) *ListNode {
	resultHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	cursor := resultHead

	for left != nil || right != nil {
		if right != nil && left == nil {
			cursor.Next = &ListNode{
				Val:  right.Val,
				Next: nil,
			}
			right = right.Next
		} else if left != nil && right == nil {
			cursor.Next = &ListNode{
				Val:  left.Val,
				Next: nil,
			}
			left = left.Next
		} else {
			if left.Val > right.Val {
				cursor.Next = &ListNode{
					Val:  right.Val,
					Next: nil,
				}
				right = right.Next
			} else {
				cursor.Next = &ListNode{
					Val:  left.Val,
					Next: nil,
				}
				left = left.Next
			}
		}
		cursor = cursor.Next
	}
	return resultHead.Next
}

func constructListNode(numbers *[]int) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cursor := head
	for _, value := range *numbers {
		cursor.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		cursor = cursor.Next
	}
	return head.Next
}
