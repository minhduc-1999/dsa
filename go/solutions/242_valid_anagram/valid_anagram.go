package validanagram

// Title: Valid Anagram

// Problem link: https://leetcode.com/problems/valid-anagram

// Testcases:
// case 0: "anagram"	"nagaram"
// case 1: "rat"	"car"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cMap := make(map[byte]uint16)
	for i := 0; i < len(s); i++ {
		_, ok := cMap[s[i]]
		if !ok {
			cMap[s[i]] = 0
		}
		cMap[s[i]] += 1
	}
	for i := 0; i < len(t); i++ {
		count, ok := cMap[t[i]]
		if !ok || count == 0 {
			return false
		}
		cMap[t[i]] -= 1
	}
	return true
}
