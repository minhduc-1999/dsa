package binarysearchtreeiterator

import "dsa/utils/stack"

// Title: Binary Search Tree Iterator

// Problem link: https://leetcode.com/problems/binary-search-tree-iterator

// Testcases:
// case 0: ["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]	[[[7,3,15,null,null,9,20]],[],[],[],[],[],[],[],[],[]]

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

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
  iter := BSTIterator{
		stack: stack,
	}
	if root == nil {
    return iter
	}
  iter.moveLeft(root)
  return iter
}

func (this *BSTIterator) moveLeft(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

func (this *BSTIterator) Next() int {
	lenS := len(this.stack)
	top := this.stack[lenS-1]
	this.stack = this.stack[:lenS-1]
	if top.Right != nil {
    this.moveLeft(top.Right)
	}
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
