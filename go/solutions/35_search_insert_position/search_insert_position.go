package searchinsertposition

// Title: Search Insert Position

// Problem link: https://leetcode.com/problems/search-insert-position

// Testcases:
// case 0: [1,3,5,6]	5
// case 1: [1,3,5,6]	2
// case 2: [1,3,5,6]	7

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		med := (right-left)/2 + left
		if nums[med] == target {
			return med
		}
		if nums[med] > target {
			right = med - 1
		} else {
			left = med + 1
		}
	}
	return left
}
