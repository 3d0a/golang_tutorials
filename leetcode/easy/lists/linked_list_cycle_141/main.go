package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	walker, runner := head, head
	for walker != nil && runner != nil && runner.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			return true
		}
	}
	return false
}

func main() {
	tail := ListNode{Val: 1, Next: nil}
	node1 := ListNode{Val: -1, Next: &tail}
	head := ListNode{Val: 4, Next: &node1}
	tail.Next = &head
	fmt.Println(hasCycle(&head))
}
