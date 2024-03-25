package binarytreelevelordertraversal

// Title: Binary Tree Level Order Traversal

// Problem link: https://leetcode.com/problems/binary-tree-level-order-traversal

// Testcases:
// case 0: [3,9,20,null,null,15,7]
// case 1: [1]
// case 2: []

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	. "dsa/utils/tree"
)

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := len(q)
		valInSameLevel := []int{}
		for i := 0; i < level; i++ {
			top := q[0]
			q = q[1:]
			valInSameLevel = append(valInSameLevel, top.Val)
			if top.Left != nil {
				q = append(q, top.Left)
			}
			if top.Right != nil {
				q = append(q, top.Right)
			}
		}
		result = append(result, valInSameLevel)
	}
	return result
}
