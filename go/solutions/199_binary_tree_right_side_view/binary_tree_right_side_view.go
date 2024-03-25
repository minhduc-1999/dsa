package binarytreerightsideview

import "dsa/utils/queue"

// Title: Binary Tree Right Side View

// Problem link: https://leetcode.com/problems/binary-tree-right-side-view

// Testcases:
// case 0: [1,2,3,null,5,null,4]
// case 1: [1,null,3]
// case 2: []

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

type NodeWithLevel struct {
	TreeNode *TreeNode
	Level    int
}

func rightSideView(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
	result := make([]int, 0, 100)
	queue := queue.NewQueue[NodeWithLevel](100)
	initNode := NodeWithLevel{
		TreeNode: root,
		Level:    0,
	}
	queue.Enqueue(initNode)
	for !queue.IsEmpty() {
		node, _ := queue.Front()
		queue.Dequeue()
    if len(result) <= node.Level {
      result = append(result, node.TreeNode.Val)
    }
		if node.TreeNode.Right != nil {
			queue.Enqueue(NodeWithLevel{
				TreeNode: node.TreeNode.Right,
				Level:    node.Level + 1,
			})
		}
		if node.TreeNode.Left != nil {
			queue.Enqueue(NodeWithLevel{
				TreeNode: node.TreeNode.Left,
				Level:    node.Level + 1,
			})
		}
	}
	return result
}
