package main
import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	resultArr := []int{}
	if root == nil {
		return resultArr
	}
	resultArr = append(resultArr, inorderTraversal(root.Left)...)
	resultArr = append(resultArr, root.Val)
	resultArr = append(resultArr, inorderTraversal(root.Right)...)
	return resultArr
}

func main() {
	root := TreeNode{Val:1}
	B := TreeNode{Val:2}
	C := TreeNode{Val:3}
	root.Right = &B
	B.Left = &C
	fmt.Println(inorderTraversal(&root))

}
