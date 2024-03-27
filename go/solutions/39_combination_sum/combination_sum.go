package combinationsum

// Title: Combination Sum

// Problem link: https://leetcode.com/problems/combination-sum

// Testcases:
// case 0: [2,3,6,7]	7
// case 1: [2,3,5]	8
// case 2: [2]	1

func backTrack(candidates []int, target int, result *[][]int, acc []int) {
	if target == 0 {
		newCombination := make([]int, len(acc))
		copy(newCombination, acc)
		*result = append(*result, newCombination)
		return
	}
	for i, cand := range candidates {
		if cand > target {
			continue
		}
		acc = append(acc, cand)
		backTrack(candidates[i:], target-cand, result, acc)
		acc = acc[:len(acc)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backTrack(candidates, target, &result, []int{})
	return result
}
