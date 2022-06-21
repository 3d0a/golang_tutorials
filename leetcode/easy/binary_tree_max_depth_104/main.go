package main
import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int{
	depth := 0
	if root == nil {
		return depth
	}
	depth += 1
	depth_left := maxDepth(root.Left)
	depth_right := maxDepth(root.Right)
	if depth_left > depth_right {
		return depth + depth_left
	}
	return depth + depth_right
}

func main() {
	root := TreeNode{Val:3}
	A := TreeNode{Val:3}
	B := TreeNode{Val:3}
  C := TreeNode{Val:3}
  D := TreeNode{Val:3}
	E := TreeNode{Val:3}

	root.Left = &A
	root.Right = &B
	B.Right = &C
	B.Left = &D
	D.Left = &E
	fmt.Println(maxDepth(&root))
}
