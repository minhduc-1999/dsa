package tree

import "dsa/utils/queue"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
		if arr[cur+1] != -9999 {
			right := &TreeNode{
				Val: arr[cur+1],
			}
			top.Right = right
			q.Enqueue(right)
		}
		cur += 2
	}
	return root
}
