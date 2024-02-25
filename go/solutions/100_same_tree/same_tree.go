package sametree

// Title: Same Tree

// Problem link: https://leetcode.com/problems/same-tree

// Testcases:
// case 0: [1,2,3]	[1,2,3]
// case 1: [1,2]	[1,null,2]
// case 2: [1,2,1]	[1,1,2]

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if q.Val != p.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
