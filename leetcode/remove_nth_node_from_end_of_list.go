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

type UserNode struct {
	userName string
	age      int
	next     *UserNode
}

func RemoteNthNodeFromEndOfList(sequence int) {
	head := constructList(10)
	fmt.Print(head)
}

func doRemove(head *UserNode, sequence int) *UserNode {
	first, second := head, head
	for index := 0; index < sequence; index++ {

	}
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
