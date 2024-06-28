package convertsortedarraytobinarysearchtree

import (
	. "dsa/utils/tree"
)

// Title: Convert Sorted Array to Binary Search Tree

// Problem link: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree

// Testcases:
// case 0: [-10,-3,0,5,9]
// case 1: [1,3]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	mid := (len(nums) - 1) / 2
	left := sortedArrayToBST(nums[0:mid])
	right := sortedArrayToBST(nums[mid+1:])
	return &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}
}
