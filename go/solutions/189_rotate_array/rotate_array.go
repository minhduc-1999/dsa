package rotatearray

// Title: Rotate Array

// Problem link: https://leetcode.com/problems/rotate-array

// Testcases:
// case 0: [1,2,3,4,5,6,7]	3
// case 1: [-1,-100,3,99]	2

func rotate(nums []int, k int) {
	k1 := k % len(nums)
	temp := make([]int, k1)
	copy(temp, nums[len(nums)-k1:])
	copy(nums[k1:], nums[:len(nums)-k1])
	copy(nums[:k1], temp)
}
