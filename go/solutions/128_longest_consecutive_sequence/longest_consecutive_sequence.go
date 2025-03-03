package longestconsecutivesequence

import "sort"

// Title: Longest Consecutive Sequence

// Problem link: https://leetcode.com/problems/longest-consecutive-sequence

// Testcases:
// case 0: [100,4,200,1,3,2]
// case 1: [0,3,7,2,5,8,4,6,0,1] 0 4 7 2 7 2 4 7 1 4
// case 2: [1,0,1,2]

func longestConsecutive(nums []int) int {
	m := map[int]int{}
	max := 0
	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			l, leftFound := m[v-1]
			if !leftFound {
				l = 0
			}
			r, rightFound := m[v+1]
			if !rightFound {
				r = 0
			}
			cur := l + r + 1
			m[v] = cur
			if leftFound {
				m[v-1] = cur
				m[v-l] = cur
			}
			if rightFound {
				m[v+1] = cur
				m[v+r] = cur
			}
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

func longestConsecutive2(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	curr := 1
	largest := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				curr++
			} else {
				largest = max(largest, curr)
				curr = 1
			}
		}
	}
	return max(largest, curr)
}
