package main

import (
	"fmt"
)

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	resultArr := []int{}
	if root == nil {
		return resultArr
	}
	resultArr = append(resultArr, root.Val)
	resultArr = append(resultArr, preorderTraversal(root.Left)...)
	resultArr = append(resultArr, preorderTraversal(root.Right)...)
	return resultArr
}

func main() {
	root := TreeNode{Val: 1}
	B := TreeNode{Val: 2}
	C := TreeNode{Val: 3}
	root.RightNode = &B
	B.LeftNode = &C
	fmt.Println(preorderTraversal(&root))

}
