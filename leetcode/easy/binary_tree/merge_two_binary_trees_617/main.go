package main
import (
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
  mainNode := &TreeNode{}
  if root1 == nil {
    return root2
  } else if root2 == nil {
    return root1
  }
  if root1 != nil && root2 != nil {
    mainNode.Val = root1.Val + root2.Val
    mainNode.Left = mergeTrees(root1.Left, root2.Left)
    mainNode.Right = mergeTrees(root1.Right, root2.Right)
  }
  return mainNode
}

func main() {

}
