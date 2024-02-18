package lettercombinationsofaphonenumber

import (
	"strconv"
)

// Title: Letter Combinations of a Phone Number

// Problem link: https://leetcode.com/problem/letter-combinations-of-a-phone-number

// Testcases:
// case 0: "23"
// case 1: ""
// case 2: "2"

func letterCombinations(digits string) []string {
	digitToLetter := map[int][]string{
		0: {""},
		1: {""},
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	result := []string{}
	letterCombinationsBackTrack(digits, 0, digitToLetter, "", &result)
	return result
}

func letterCombinationsBackTrack(digits string, curIndex int, digitToLetter map[int][]string, accumulator string, result *[]string) {
	if curIndex >= len(digits) {
		if accumulator != "" {
			*result = append(*result, accumulator)
		}
		return
	}
	digit, _ := strconv.ParseInt(string(digits[curIndex]), 10, 0)
	letters := digitToLetter[int(digit)]
	for _, letter := range letters {
		accumulator = accumulator + letter
		letterCombinationsBackTrack(digits, curIndex+1, digitToLetter, accumulator, result)
		accumulator = accumulator[:len(accumulator)-1]
	}
}
