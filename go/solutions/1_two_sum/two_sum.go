package twosum

// Title: Two Sum

// Problem link: https://leetcode.com/problems/two-sum

// Testcases:
// case 0: [2,7,11,15]	9
// case 1: [3,2,4]	6
// case 2: [3,3]	6

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for index, value := range nums {
		indexMap[value] = index
	}
	for i, value := range nums {
		foundIndex, ok := indexMap[target-value]
		if ok && foundIndex != i {
			return []int{i, foundIndex}
		}
	}
	return []int{}
}
