package isomorphicstrings

// Title: Isomorphic Strings

// Problem link: https://leetcode.com/problems/isomorphic-strings

// Testcases:
// case 0: "egg"	"add"
// case 1: "foo"	"bar"
// case 2: "paper"	"title"

func isIsomorphic(s string, t string) bool {
  sLen := len(s)
  sMap := map[byte]byte{}
  tMap := map[byte]byte{}
  for i := 0; i < sLen; i++ {
    sValue, sOk := sMap[s[i]]
    tValue, tOk := tMap[t[i]]
    if !sOk && !tOk {
      sMap[s[i]] = t[i]
      tMap[t[i]] = s[i]
      continue
    }
    if sValue != t[i] || tValue != s[i] {
      return false
    }
  }
  return true
}
