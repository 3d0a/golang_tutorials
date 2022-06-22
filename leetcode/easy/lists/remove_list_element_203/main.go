package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var stepBack, resultHead *ListNode
	if head == nil {
		return nil
	}
	for head != nil {
		if head.Val != val {
			stepBack = head
			resultHead = stepBack
			head = head.Next
			break
		} else {
			head = head.Next
		}
		if head == nil {
			return nil
		}
	}
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			stepBack.Next = head
			stepBack = stepBack.Next
			head = head.Next
		}
		if head == nil {
			stepBack.Next = nil
			break
		}
	}
	return resultHead
}

func main() {
	A := ListNode{Val: 1}
	B := ListNode{Val: 6}
	C := ListNode{Val: 3}
	D := ListNode{Val: 1}
	E := ListNode{Val: 5}
	F := ListNode{Val: 7}
	G := ListNode{Val: 1}

	A.Next = &B
	B.Next = &C
	C.Next = &D
	D.Next = &E
	E.Next = &F
	F.Next = &G
	G.Next = nil

	list := removeElements(&A, 1)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
