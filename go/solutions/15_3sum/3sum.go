package threesum

import "sort"

// Title: 3Sum

// Problem link: https://leetcode.com/problems/3sum

// Testcases:
// case 0: [-1,0,1,2,-1,-4]
// case 1: [0,1,1]
// case 2: [0,0,0]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i, v := range nums {
		if i != 0 && v == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if l-1 != i && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r != len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if nums[l]+nums[r] == -v {
				result = append(result, []int{v, nums[l], nums[r]})
				l++
				r--
				continue
			}
			if nums[l]+nums[r] < -v {
				l++
				continue
			}
			r--
			continue
		}
	}
	return result
}
