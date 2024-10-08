package minimumsizesubarraysum

// Title: Minimum Size Subarray Sum

// Problem link: https://leetcode.com/problems/minimum-size-subarray-sum

// Testcases:
// case 0: 7	[2,3,1,2,4,3]
// case 1: 4	[1,4,4]
// case 2: 11	[1,1,1,1,1,1,1,1]

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	minLen, l, r, sum := n+1, 0, 0, 0
	for r < n && l <= r {
		sum += nums[r]
		if sum >= target {
			curLen := r - l + 1
			if curLen < minLen {
				minLen = curLen
			}
			sum -= (nums[l] + nums[r])
			l += 1
		} else {
			r += 1
		}
	}
	if minLen == n+1 {
		return 0
	}
	return minLen
}
