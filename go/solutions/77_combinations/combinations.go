package combinations

// Title: Combinations

// Problem link: https://leetcode.com/problems/combinations

// Testcases:
// case 0: 4	2
// case 1: 1	1

func combine(n int, k int) [][]int {
	result := [][]int{}
	accumulator := []int{}
	backTrack(1, n, k, accumulator, &result)
	return result
}

func backTrack(i int, n int, k int, accummulator []int, result *[][]int) {
	if k == 0 {
		temp := make([]int, len(accummulator))
		copy(temp, accummulator)
		*result = append(*result, temp)
		return
	}
  if i > n {
    return
  }
	for j := i; j <= n; j++ {
		accummulator = append(accummulator, j)
		backTrack(j+1, n, k-1, accummulator, result)
		accummulator = accummulator[:len(accummulator)-1]
	}
}
