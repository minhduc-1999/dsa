package pathsum

// Title: Path Sum

// Problem link: https://leetcode.com/problems/path-sum

// Testcases:
// case 0: [5,4,8,11,null,13,4,7,2,null,null,null,1]	22
// case 1: [1,2,3]	5
// case 2: []	0

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

func hasPathSum(root *TreeNode, targetSum int) bool {
  if root == nil {
    return false
  }
  if root.Left == nil && root.Right == nil {
    return root.Val == targetSum
  }

  return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}
