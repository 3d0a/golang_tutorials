package main
import (
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  currentList := &ListNode{}
  head := currentList
  saved := 0
  for l1 != nil || l2 != nil {
    if l1 == nil {
       currentList.Val += l2.Val
    } else if l2 == nil {
      currentList.Val += + l1.Val
    } else {
      currentList.Val += l1.Val + l2.Val
    }
    if currentList.Val >= 10 {
      currentList.Val = currentList.Val - 10
      saved = 1
    } else {
      saved = 0
    }
    if l1 != nil {l1 = l1.Next}
    if l2 != nil {l2 = l2.Next}
    if l1 != nil || l2 != nil || saved > 0 {
      currentList.Next = &ListNode{Val:saved}
      currentList = currentList.Next
    }
  }
  return head
}

func main() {
	node1_1 := ListNode{Val:1, Next:nil}
	node1_2 := ListNode{Val:4, Next:&node1_1}
	node1_3 := ListNode{Val:4, Next:&node1_2}
	node1_4 := ListNode{Val:4, Next:&node1_3}

	node2_1 := ListNode{Val:1, Next:nil}
  node2_2 := ListNode{Val:4, Next:&node2_1}
  node2_3 := ListNode{Val:5, Next:&node2_2}
	addTwoNumbers(&node1_4, &node2_3)
}
