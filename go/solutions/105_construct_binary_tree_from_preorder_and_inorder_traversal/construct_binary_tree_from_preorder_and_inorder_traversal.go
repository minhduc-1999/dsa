package constructbinarytreefrompreorderandinordertraversal

// Title: Construct Binary Tree from Preorder and Inorder Traversal

// Problem link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

// Testcases:
// case 0: [3,9,20,15,7]	[9,3,15,20,7]
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
func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 || len(inorder) == 0 {
    return nil
  }
	rootVal := preorder[0]
	rootInorderIdx := findIndex(inorder, rootVal)

	leftNode := buildTree(preorder[1:1+rootInorderIdx], inorder[:rootInorderIdx])
	rightNode := buildTree(preorder[1+rootInorderIdx:], inorder[rootInorderIdx+1:])
	return &TreeNode{
		Val:   rootVal,
		Left:  leftNode,
		Right: rightNode,
	}
}
