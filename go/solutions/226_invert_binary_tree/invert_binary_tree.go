package invertbinarytree

// Title: Invert Binary Tree

// Problem link: https://leetcode.com/problems/invert-binary-tree

// Testcases:
// case 0: [4,2,7,1,3,6,9]
// case 1: [2,1,3]
// case 2: []

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
  if root != nil {
    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
  }
  return root
}
