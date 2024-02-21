package permutations

// Title: Permutations

// Problem link: https://leetcode.com/problems/permutations

// Testcases:
// case 0: [1,2,3]
// case 1: [0,1]
// case 2: [1]

func permute(nums []int) [][]int {
	result := [][]int{}
	permuteBackTrack(nums, 0, len(nums)-1, &result)
	return result
}

func swap(nums []int, lhs int, rhs int) {
	temp := nums[lhs]
	nums[lhs] = nums[rhs]
	nums[rhs] = temp
}

func permuteBackTrack(nums []int, left int, right int, result *[][]int) {
	if left == right {
    per := make([]int, len(nums))
    copy(per, nums)
    *result = append(*result, per) 
		return
	}
	for i := left; i <= right; i++ {
		swap(nums, left, i)
		permuteBackTrack(nums, left+1, right, result)
		swap(nums, left, i)
	}
}
