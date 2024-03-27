package generateparentheses

import "strings"

// Title: Generate Parentheses

// Problem link: https://leetcode.com/problems/generate-parentheses

// Testcases:
// case 0: 3
// case 1: 1

func backTrack(nOpen int, nClose int, result *[]string, acc string) {
  if nOpen == 0 && nClose == 0 {
    temp := strings.Clone(acc)
    *result = append(*result, temp)
    return
  }
  if nOpen > 0 {
    acc += "("
    backTrack(nOpen - 1, nClose, result, acc)
    acc = acc[:len(acc) - 1]
  }
  if nClose  > nOpen {
    acc += ")"
    backTrack(nOpen, nClose - 1, result, acc)
    acc = acc[:len(acc) - 1]
  }
}
func generateParenthesis(n int) []string {
  result := []string{}
  backTrack(n, n, &result, "")
  return result
}
