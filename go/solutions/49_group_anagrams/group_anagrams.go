package groupanagrams

import (
	"sort"
)

// Title: Group Anagrams

// Problem link: https://leetcode.com/problems/group-anagrams

// Testcases:
// case 0: ["eat","tea","tan","ate","nat","bat"]
// case 1: [""]
// case 2: ["a"]

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(s string) string {
	r := sortRunes([]rune(s))
	sort.Sort(r)
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		sortedKey := sortString(str)
		arr, ok := m[sortedKey]
		if !ok {
			m[sortedKey] = []string{str}
		} else {
			m[sortedKey] = append(arr, str)
		}
	}
	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
