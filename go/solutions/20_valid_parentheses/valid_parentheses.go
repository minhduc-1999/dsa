package validparentheses

import "dsa/utils/stack"
// Title: Valid Parentheses

// Problem link: https://leetcode.com/problems/valid-parentheses

// Testcases:
// case 0: "()"
// case 1: "()[]{}"
// case 2: "(]"

func isValid(s string) bool {
  openPair := map[rune]rune{
    '}': '{',
    ']':'[',
    ')':'(',
  }
  stack := stack.NewStack[rune](len(s))
  for _, ch := range s {
    if ch == '(' || ch == '[' || ch == '{' {
      stack.Push(ch)
      continue
    }

    openCh, err := stack.Top()
    stack.Pop()
    if err != nil {
      return false
    }
    if openCh != openPair[ch] {
      return false
    }
  }
  return stack.IsEmpty()
}
