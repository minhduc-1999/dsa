package containerwithmostwater

// Title: Container With Most Water

// Problem link: https://leetcode.com/problems/container-with-most-water

// Testcases:
// case 0: [1,8,6,2,5,4,8,3,7]
// case 1: [1,1]

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	water := 0
	for l < r {
		newWater := min(height[l], height[r]) * (r - l)
		if newWater > water {
			water = newWater
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return water
}
