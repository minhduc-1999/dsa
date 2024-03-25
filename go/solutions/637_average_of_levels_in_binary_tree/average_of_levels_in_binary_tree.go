package averageoflevelsinbinarytree

// Title: Average of Levels in Binary Tree

// Problem link: https://leetcode.com/problems/average-of-levels-in-binary-tree

// Testcases:
// case 0: [3,9,20,null,null,15,7]
// case 1: [3,9,20,15,7]

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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	q := []*TreeNode{}
	result := []float64{}
	q = append(q, root)

	for len(q) > 0 {
		level := len(q)
		sum := 0
		for i := 0; i < level; i++ {
			top := q[0]
			q = q[1:]
			sum += top.Val
			if top.Left != nil {
				q = append(q, top.Left)
			}
			if top.Right != nil {
				q = append(q, top.Right)
			}
		}
		result = append(result, float64(sum)/float64(level))
	}
	return result
}
