package leetcode

import (
	"fmt"
	"strconv"
)

//19. Remove Nth Node From End of List
//Medium
//Given a linked list, remove the n-th node from the end of list and return its head.
//
//Example:
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//
//Given n will always be valid.
//
//Follow up:
//
//Could you do this in one pass?
//基本处理链表的的节点问题，如果是正顺序，那么很好处理，就是直接遍历就行
//但是如果是倒序处理的，那就要至少两个指针了，或者是，你可以首先翻转链表，转化为正序
//之后再次翻转链表就行，只是o(n)多个了常数倍而已
type UserNode struct {
	userName string
	age      int
	next     *UserNode
}

func RemoteNthNodeFromEndOfList(sequence int) {
	head := constructList(10)
	doRemove(head, 3)
	fmt.Print(head)
}

func doRemove(head *UserNode, sequence int) *UserNode {
	first, second := head, head
	for index := 0; index < sequence; index++ {
		first = first.next
	}
	for first.next != nil {
		second = second.next
		first = first.next
	}
	second.next = second.next.next
	return head
}

func constructList(length int) *UserNode {
	var head, temp *UserNode
	for index := 0; index < length; index++ {
		if head == nil {
			head = &UserNode{userName: "userName" + strconv.Itoa(index), age: index}
			temp = head
			continue
		}
		if temp != nil {
			temp.next = &UserNode{userName: "userName" + strconv.Itoa(index), age: index}
			temp = temp.next
		}
	}
	return head
}
