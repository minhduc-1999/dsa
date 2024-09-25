package containsduplicateii

import "math"

// Title: Contains Duplicate II

// Problem link: https://leetcode.com/problems/contains-duplicate-ii

// Testcases:
// case 0: [1,2,3,1]	3
// case 1: [1,0,1,1]	1
// case 2: [1,2,3,1,2,3]	2

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i, value := range nums {
		num, ok := indexMap[value]
		if ok && int(math.Abs(float64(num-i))) <= k {
			return true
		}
		indexMap[value] = i
	}
	return false
}
