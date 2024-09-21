package wordpattern

import "strings"

// Title: Word Pattern

// Problem link: https://leetcode.com/problems/word-pattern

// Testcases:
// case 0: "abba"	"dog cat cat dog"
// case 1: "abba"	"dog cat cat fish"
// case 2: "aaaa"	"dog cat cat dog"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	patternMap := make(map[string]byte)
	wordMap := make(map[byte]string)
	for i, word := range words {
		foundChar, ok := patternMap[word]
		if ok {
			if foundChar != pattern[i] {
				return false
			}
		} else {
			patternMap[word] = pattern[i]
		}
		foundWord, ok := wordMap[pattern[i]]
		if ok {
			if foundWord != word {
				return false
			}
		} else {
			wordMap[pattern[i]] = word
		}
	}
	return true
}
