package main

import (
	"fmt"
)

type ListNode struct {
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	depth := 0
	for {
		if current == nil {
			break
		}
		depth++
		current = current.Next
	}
	nodeToDelete := depth - n + 1
	depth = 1
	current = head
	if nodeToDelete == 1 {
		return current.Next
	}
	for {
		if nodeToDelete == depth+1 {
			current.Next = current.Next.Next
			break
		}
		depth++
		current = current.Next
	}
	return head
}

func main() {

}
