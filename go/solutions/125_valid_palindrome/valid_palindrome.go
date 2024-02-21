package validpalindrome

import (
	"strings"
)

// Title: Valid Palindrome

// Problem link: https://leetcode.com/problems/valid-palindrome

// Testcases:
// case 0: "A man, a plan, a canal: Panama"
// case 1: "race a car"
// case 2: " "

func isAlphanumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
