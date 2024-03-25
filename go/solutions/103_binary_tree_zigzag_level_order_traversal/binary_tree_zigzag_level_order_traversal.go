package binarytreezigzaglevelordertraversal

import . "dsa/utils/tree"

// Title: Binary Tree Zigzag Level Order Traversal

// Problem link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	direction := true
	for len(q) > 0 {
		level := len(q)
		valueInSameLevel := []int{}
		temp := make([]*TreeNode, level)
		for i := 0; i < level; i++ {
			top := q[0]
			q = q[1:]
			valueInSameLevel = append(valueInSameLevel, top.Val)
			temp[i] = top
		}
		for i := len(temp) - 1; i >= 0; i-- {
			node := temp[i]
			if direction {
				if node.Right != nil {
					q = append(q, node.Right)
				}
				if node.Left != nil {
					q = append(q, node.Left)
				}
			} else {
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
		}
		result = append(result, valueInSameLevel)
		direction = !direction
	}
	return result
}
