package trappingrainwater

// Title: Trapping Rain Water

// Problem link: https://leetcode.com/problems/trapping-rain-water

// Testcases:
// case 0: [0,1,0,2,1,0,1,3,2,1,2,1]
// case 1: [4,2,0,3,2,5]

func trap(height []int) int {
	l, r := 0, len(height)-1
	water := 0
	i := 1
	for i > l && i < r {
		if height[l] < height[r] {
			i = l + 1
			for i < r {
				if height[i] >= height[l] {
					l = i
					i++
					break
				} else {
					water += height[l] - height[i]
					i++
				}
			}
		} else {
			i = r - 1
			for i > l {
				if height[i] >= height[r] {
					r = i
					i--
					break
				} else {
					water += height[r] - height[i]
					i--
				}
			}
		}
	}
	return water
}
