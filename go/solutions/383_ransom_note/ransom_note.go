package ransomnote

// Title: Ransom Note

// Problem link: https://leetcode.com/problems/ransom-note

// Testcases:
// case 0: "a"	"b"
// case 1: "aa"	"ab"
// case 2: "aa"	"aab"

func canConstruct(ransomNote string, magazine string) bool {
  if len(ransomNote) > len(magazine) {
    return false
  }
  mgHash := make([]int, 26) // hashmap of 26 alphabet letters
  for _, c := range magazine {
    mgHash[c - 'a']++
  }
  for _, c := range ransomNote {
    if mgHash[c - 'a'] == 0 {
      return false
    }
    mgHash[c - 'a']--
  }
  return true
}
