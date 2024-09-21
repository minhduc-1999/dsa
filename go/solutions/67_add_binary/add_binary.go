package addbinary

import "fmt"

// Title: Add Binary

// Problem link: https://leetcode.com/problems/add-binary

// Testcases:
// case 0: "11"	"1"
// case 1: "1010"	"1011"

func addBinary(a string, b string) string {
	i1, i2 := len(a)-1, len(b)-1
	result := ""
	var carry byte
	for i1 >= 0 || i2 >= 0 || carry > 0 {
		if i1 >= 0 {
			carry += a[i1] - '0'
			i1--
		}
		if i2 >= 0 {
			carry += b[i2] - '0'
			i2--
		}
		result = fmt.Sprintf("%v", carry%2) + result
		carry /= 2
	}
	return result
}
