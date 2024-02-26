package populatingnextrightpointersineachnodeii

import "dsa/utils/queue"

// Title: Populating Next Right Pointers in Each Node II

// Problem link: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii

// Testcases:
// case 0: [1,2,3,4,5,null,7]
// case 1: []

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := queue.NewQueue[*Node](6000)
	queue.Enqueue(root)
	queue.Enqueue(nil)
  var current *Node
	for !queue.IsEmpty() {
		current, _ = queue.Front()
		queue.Dequeue()
    if current == nil {
      continue
    }

    nextNode, _:= queue.Front()
    current.Next = nextNode
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
		if current.Next == nil {
			queue.Enqueue(nil)
		}
	}
	return root
}
