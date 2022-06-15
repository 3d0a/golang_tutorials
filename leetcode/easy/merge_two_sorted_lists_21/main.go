package main
import (
	"fmt"
)


type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  finalList := &ListNode{}
  headFinalList := finalList
  if list1 == nil && list2 == nil {
    return nil
  } else if list1 == nil && list2 != nil {
    return list2
  } else if list2 == nil && list1 != nil {
    return list1
  }
  for  {
      if list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
          finalList.Val = list1.Val
          list1 = list1.Next
        } else if list1.Val >= list2.Val {
          finalList.Val = list2.Val
          list2 = list2.Next
        }
      }
      if list1 != nil && list2 != nil {
        fmt.Println(finalList.Val)
        finalList.Next = &ListNode{}
        finalList = finalList.Next
      } else if list1 == nil && list2 == nil {
        return nil
      } else  {
        if list1 == nil {
          finalList.Next = list2
          break
        } else if list2 == nil {
          finalList.Next = list1
          break
        }
    }
  }
  return headFinalList
}

func main() {
	tail1 :=   ListNode{Val:4, Next:nil}
	list1_2 := ListNode{Val:2, Next:&tail1}
	list1_3 := ListNode{Val:1, Next:&list1_2}

  tail2 :=   ListNode{Val:4, Next:nil}
  list2_2 := ListNode{Val:3, Next:&tail2}
  list2_3 := ListNode{Val:1, Next:&list2_2}
	list := mergeTwoLists(&list1_3, &list2_3)
	for list.Next != nil {
		list = list.Next
		fmt.Println(list.Val)
	}

}
