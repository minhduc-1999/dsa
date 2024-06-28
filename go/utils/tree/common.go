package tree

import (
	"dsa/utils/queue"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTreeBFS(root *TreeNode) string {
	s := ""
	if root == nil {
		return s
	}
	q := queue.NewQueue[*TreeNode](100)
	q.Enqueue(root)
	for !q.IsEmpty() {
		top, _ := q.Front()
		q.Dequeue()
		if top.Left != nil {
			q.Enqueue(top.Left)
		}
		if top.Right != nil {
			q.Enqueue(top.Right)
		}
		if top.Val == -9999 {
			s += "null\t"
		} else {
			s += fmt.Sprintf("%d\t", top.Val)
		}
	}
	return s
}

func NewFromArray(arr []int) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: arr[0],
	}
	q := queue.NewQueue[*TreeNode](n)
	q.Enqueue(root)
	cur := 1
	for cur < n {
		top, _ := q.Front()
		q.Dequeue()
		if arr[cur] != -9999 {
			left := &TreeNode{
				Val: arr[cur],
			}
			top.Left = left
			q.Enqueue(left)
		}
		cur += 1
		if cur >= n {
			break
		}
		if arr[cur] != -9999 {
			right := &TreeNode{
				Val: arr[cur],
			}
			top.Right = right
			q.Enqueue(right)
		}
		cur += 1
	}
	return root
}
