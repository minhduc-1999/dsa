package issubsequence

// Title: Is Subsequence

// Problem link: https://leetcode.com/problems/is-subsequence

// Testcases:
// case 0: "abc"	"ahbgdc"
// case 1: "axc"	"ahbgdc"

func isSubsequence(s string, t string) bool {
  i, j := 0, 0
  lenS, lenT := len(s), len(t)
  for {
    if i == lenS {
      return true
    }
    if j == lenT {
      return false
    }
    if s[i] == t[j] {
      i++
      j++
      continue
    }
    j++
  }
}
