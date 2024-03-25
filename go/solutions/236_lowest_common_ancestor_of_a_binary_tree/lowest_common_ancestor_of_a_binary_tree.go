package lowestcommonancestorofabinarytree

// Title: Lowest Common Ancestor of a Binary Tree

// Problem link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

// Testcases:
// case 0: [3,5,1,6,2,0,8,null,null,7,4]	5	1
// case 1: [3,5,1,6,2,0,8,null,null,7,4]	5	4
// case 2: [1,2]	1	2

/* *
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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	leftLca := lowestCommonAncestor(root.Left, p, q)
	rightLca := lowestCommonAncestor(root.Right, p, q)
	if leftLca == nil && rightLca == nil {
		return nil
	}
	if leftLca == nil {
		return rightLca
	}
	if rightLca == nil {
		return leftLca
	}
	return root
}
