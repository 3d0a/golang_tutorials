package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	depth := 0
	current := head
	for {
		if current == nil {
			break
		}
		depth++
		current = current.Next
	}
	half := depth / 2 + 1
	depth = 1
	current = head
	for {
		if depth == half {
			return current
		}
		depth ++
		current = current.Next
	}
	return nil
}

func main() {

}
