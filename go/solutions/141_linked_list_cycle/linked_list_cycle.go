package linkedlistcycle

// Title: Linked List Cycle

// Problem link: https://leetcode.com/problems/linked-list-cycle

// Testcases:
// case 0: [3,2,0,-4]	1
// case 1: [1,2]	0
// case 2: [1]	-1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
  fast, slow := head, head
  for fast != nil && fast.Next != nil {
    fast = fast.Next.Next
    slow = slow.Next
    if fast == slow {
      return true
    }
  }
	return false
}
