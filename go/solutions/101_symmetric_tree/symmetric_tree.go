package symmetrictree

// Title: Symmetric Tree

// Problem link: https://leetcode.com/problems/symmetric-tree

// Testcases:
// case 0: [1,2,2,3,4,4,3]
// case 1: [1,2,2,null,3,null,3]

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

func checkSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}
