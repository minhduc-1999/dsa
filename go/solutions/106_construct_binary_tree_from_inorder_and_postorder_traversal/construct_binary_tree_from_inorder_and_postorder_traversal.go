package constructbinarytreefrominorderandpostordertraversal

// Title: Construct Binary Tree from Inorder and Postorder Traversal

// Problem link: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal

// Testcases:
// case 0: [9,3,15,20,7]	[9,15,7,20,3]
// case 1: [-1]	[-1]

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

func findIndex(arr []int, item int) int {
	for i, elem := range arr {
		if elem == item {
			return i
		}
	}
	return -1
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	rootInorderIdx := findIndex(inorder, rootVal)

	leftNode := buildTree(inorder[:rootInorderIdx], postorder[:rootInorderIdx])
	rightNode := buildTree(inorder[rootInorderIdx+1:], postorder[rootInorderIdx:len(postorder)-1])
	return &TreeNode{
		Val:   rootVal,
		Left:  leftNode,
		Right: rightNode,
	}
}
