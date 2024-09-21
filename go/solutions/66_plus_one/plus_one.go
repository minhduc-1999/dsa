package plusone

// Title: Plus One

// Problem link: https://leetcode.com/problems/plus-one

// Testcases:
// case 0: [1,2,3]
// case 1: [4,3,2,1]
// case 2: [9]

func plusOne(digits []int) []int {
	i := len(digits) - 1
	carry := 1
	for i >= 0 {
		carry += digits[i]
		if carry > 9 {
			digits[i] = carry - 10
			carry = 1
		} else {
			digits[i] = carry
			carry = 0
			break
		}
		i--
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
