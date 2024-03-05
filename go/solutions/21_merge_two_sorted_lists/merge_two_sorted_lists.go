package mergetwosortedlists

// Title: Merge Two Sorted Lists

// Problem link: https://leetcode.com/problems/merge-two-sorted-lists

// Testcases:
// case 0: [1,2,4]	[1,3,4]
// case 1: []	[]
// case 2: []	[0]

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
    cur = cur.Next
	}
	return head.Next
}
