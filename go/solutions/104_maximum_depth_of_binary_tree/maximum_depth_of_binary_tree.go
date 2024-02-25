package maximumdepthofbinarytree

// Title: Maximum Depth of Binary Tree

// Problem link: https://leetcode.com/problems/maximum-depth-of-binary-tree

// Testcases:
// case 0: [3,9,20,null,null,15,7]
// case 1: [1,null,2]

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := 1 + maxDepth(root.Left)
	rDepth := 1 + maxDepth(root.Right)
	return max(lDepth, rDepth)
}
