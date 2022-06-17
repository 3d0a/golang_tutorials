package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
  if head.Next == nil {
    return head
  }
	prev := head
	currentHead := head.Next
	prev.Next = nil
	for {
		tmp := currentHead.Next
		if currentHead.Next == nil {
			currentHead.Next = prev
			break
		}
		currentHead.Next = prev
		prev = currentHead
		currentHead = tmp
	}
	return currentHead
}

func main() {

	A := ListNode{Val:1}
	B := ListNode{Val:2}
	C := ListNode{Val:3}
  D := ListNode{Val:4}
  E := ListNode{Val:5}
  F := ListNode{Val:6}

	A.Next = &B
	B.Next = &C
  C.Next = &D
  D.Next = &E
  E.Next = &F
  F.Next = nil
	list := reverseList(&A)

	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}


