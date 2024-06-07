package minimumabsolutedifferenceinbst

import (
	. "dsa/utils/tree"
	"math"
)

// Title: Minimum Absolute Difference in BST

// Problem link: https://leetcode.com/problems/minimum-absolute-difference-in-bst

// Testcases:
// case 0: [4,2,6,1,3]
// case 1: [1,0,48,null,null,12,49]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leftMostNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root
	for temp.Left != nil {
		temp = temp.Left
	}
	return temp
}
func rightMostNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp
}
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return -1
	}
	min := math.MaxInt
	if root.Left != nil {
		rMost := rightMostNode(root.Left)
		lMin := root.Val - rMost.Val
		lSubMin := getMinimumDifference(root.Left)
		if min > lSubMin {
			min = lSubMin
		}
		if min > lMin {
			min = lMin
		}
	}
	if root.Right != nil {
		lMost := leftMostNode(root.Right)
		rMin := lMost.Val - root.Val
		rSubMin := getMinimumDifference(root.Right)
		if min > rMin {
			min = rMin
		}
		if min > rSubMin {
			min = rSubMin
		}
	}
	return min
}
