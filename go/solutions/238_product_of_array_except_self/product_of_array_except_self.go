package productofarrayexceptself

// Title: Product of Array Except Self

// Problem link: https://leetcode.com/problems/product-of-array-except-self

// Testcases:
// case 0: [1,2,3,4]
// 24 12 8 6
// 1 2 3 4
// 1 1 2 6
// 24 12 4 1
// case 1: [-1,1,0,-3,3]

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1
	i := 1
	for i < len(nums) {
		result[i] = result[i-1] * nums[i-1]
		i++
	}
	i = len(nums) - 2
	lastP := 1
	for i > -1 {
		result[i] *= nums[i+1] * lastP
		lastP = nums[i+1] * lastP
		i--
	}
	return result
}
