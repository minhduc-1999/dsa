package twosumiiinputarrayissorted

// Title: Two Sum II - Input Array Is Sorted

// Problem link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

// Testcases:
// case 0: [2,7,11,15]	9
// case 1: [2,3,4]	6
// case 2: [-1,0]	-1

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	sum := numbers[l] + numbers[r]
	for l < r {
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
			sum += numbers[l] - numbers[l-1]
		} else {
			r--
			sum += numbers[r] - numbers[r+1]
		}
	}
	return []int{l + 1, r + 1}
}
