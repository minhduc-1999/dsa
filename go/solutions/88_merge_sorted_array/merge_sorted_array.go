package mergesortedarray

// Title: Merge Sorted Array

// Problem link: https://leetcode.com/problems/merge-sorted-array

// Testcases:
// case 0: [1,2,3,0,0,0]	3	[2,5,6]	3
// case 1: [1]	1	[]	0
// case 2: [0]	0	[1]	1

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, i1, i2 := len(nums1)-1, m-1, len(nums2)-1
	for i1 >= 0 && i2 >= 0 {
		if nums2[i2] >= nums1[i1] {
			nums1[i] = nums2[i2]
			i2--
		} else {
			nums1[i] = nums1[i1]
			i1--
		}
		i--
	}
	if i1 < 0 {
		for i2 >= 0 {
			nums1[i] = nums2[i2]
			i2--
			i--
		}
	}
}
