package sumroottoleafnumbers

import (
	"fmt"
	"strconv"
)

// Title: Sum Root to Leaf Numbers

// Problem link: https://leetcode.com/problems/sum-root-to-leaf-numbers

// Testcases:
// case 0: [1,2,3]
// case 1: [4,9,0,5,1]

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

func buildNumbers(node *TreeNode, num string, result *[]int) {
	if node == nil {
		return
	}
	num = num + fmt.Sprint(node.Val)
	if node.Left == nil && node.Right == nil {
		temp, _ := strconv.ParseInt(num, 10, 32)
		*result = append(*result, int(temp))
		return
	}
	buildNumbers(node.Left, num, result)
	buildNumbers(node.Right, num, result)
}

func sumNumbers(root *TreeNode) int {
	numbers := []int{}
	buildNumbers(root, "", &numbers)
	sum := 0
	for _, elem := range numbers {
		sum += elem
	}
	return sum
}
