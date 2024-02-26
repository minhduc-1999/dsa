package countcompletetreenodes

// Title: Count Complete Tree Nodes

// Problem link: https://leetcode.com/problems/count-complete-tree-nodes

// Testcases:
// case 0: [1,2,3,4,5,6]
// case 1: []
// case 2: [1]

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

func Recursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Recursion(root.Left) + Recursion(root.Right)
}

func GetMaxLeftIndex(root *TreeNode, idx int) int {
	if root == nil {
		return -1
	}
	for root.Left != nil {
		root = root.Left
		idx = 2*idx + 1
	}
	return idx
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := root
	curIdx := 0
	maxLeftIndex := GetMaxLeftIndex(cur, curIdx)
	for cur != nil {
		maxLeftRightIndex := GetMaxLeftIndex(cur.Right, curIdx*2+2)
		if maxLeftIndex == maxLeftRightIndex || maxLeftRightIndex < 0 {
			return maxLeftIndex + 1
		}
		if maxLeftIndex < maxLeftRightIndex {
			cur = cur.Right
			curIdx = curIdx*2 + 2
			maxLeftIndex = maxLeftRightIndex
		} else {
			cur = cur.Left
			curIdx = curIdx*2 + 1
		}
	}
	return maxLeftIndex
}
