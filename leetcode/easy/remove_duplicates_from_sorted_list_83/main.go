package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	current := head.Next
	prev    := head
	toReturn := prev
	for {
		if current == nil {
			prev.Next = nil 
			break
		}
		if current.Val == prev.Val {
			current = current.Next
		} else {
			prev.Next = current
			prev = prev.Next
			current = current.Next
		}
	}
	return toReturn
}

func main() {
	A := ListNode{Val:1}
	B := ListNode{Val:2}
  C := ListNode{Val:3}
  D := ListNode{Val:3}
  E := ListNode{Val:4}
  F := ListNode{Val:4}
  G := ListNode{Val:4}

	A.Next = &B
	B.Next = &C
	C.Next = &D
  D.Next = &E
  E.Next = &F
  F.Next = &G
  G.Next = nil

	list := deleteDuplicates(&A)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
