package happynumber

import "math"

// Title: Happy Number

// Problem link: https://leetcode.com/problems/happy-number

// Testcases:
// case 0: 19
// case 1: 2

func isHappy(n int) bool {
	visited := make(map[int]struct{})
	sum := 0
	for {
		if n == 0 {
			if sum == 1 {
				return true
			}
			_, found := visited[sum]
			if found {
				return false
			}
			visited[sum] = struct{}{}
			n = sum
			sum = 0
		}
		digit := n % 10
		n /= 10
		sum += int(math.Pow(float64(digit), 2))
	}
}
