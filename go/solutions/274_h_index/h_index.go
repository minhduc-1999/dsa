package hindex

import (
	"slices"
)

// Title: H-Index

// Problem link: https://leetcode.com/problems/h-index

// Testcases:
// case 0: [3,0,6,1,5]
// case 1: [1,3,1]

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(i, j int) int { return j - i })
	hIndex := len(citations)
	i := 0
	for i < len(citations) {
		if i == hIndex {
			break
		}
		if citations[i] < hIndex {
			hIndex--
			i = 0
		} else {
			i++
		}
	}
	return hIndex
}
