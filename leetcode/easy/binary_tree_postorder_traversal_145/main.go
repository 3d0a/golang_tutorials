package main
import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int  {
	resultArr := []int{}
	if root == nil {
		return resultArr
	}
	resultArr = append(resultArr, postorderTraversal(root.Left)...)
	resultArr = append(resultArr, postorderTraversal(root.Right)...)
	resultArr = append(resultArr, root.Val)
	return resultArr
}

func main() {
	root := TreeNode{Val:1}
	B := TreeNode{Val:2}
	C := TreeNode{Val:3}
	root.Right = &B
	B.Left = &C
	fmt.Println(postorderTraversal(&root))
}
