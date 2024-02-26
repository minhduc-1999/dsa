package tree

import (
  . "dsa/utils/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(n int) *TreeNode {
	root := TreeNode{
		Val: 1,
	}
	q := NewQueue[*TreeNode](n)
	q.Enqueue(&root)
	i := 1
	for i <= n {
		cur, _ := q.Front()
		q.Dequeue()
		cur.Left = &TreeNode{
			Val: i + 1,
		}
		q.Enqueue(cur.Left)
		cur.Right = &TreeNode{
			Val: i + 2,
		}
		q.Enqueue(cur.Right)
    i = i + 2
	}
	return &root
}

