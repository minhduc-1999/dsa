package flattenbinarytreetolinkedlist

import "dsa/utils/stack"

// Title: Flatten Binary Tree to Linked List

// Problem link: https://leetcode.com/problems/flatten-binary-tree-to-linked-list

// Testcases:
// case 0: [1,2,5,3,4,null,6]
// case 1: []
// case 2: [0]

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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := stack.NewStack[*TreeNode](2000)
	stack.Push(root)
	var cur *TreeNode
	for !stack.IsEmpty() {
		cur, _ = stack.Top()
		stack.Pop()
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
		next, err := stack.Top()
		if err == nil {
			cur.Right = next
		}
    cur.Left = nil
	}
}
