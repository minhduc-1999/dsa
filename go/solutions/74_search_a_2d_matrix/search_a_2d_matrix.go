package searcha2dmatrix

// Title: Search a 2D Matrix

// Problem link: https://leetcode.com/problems/search-a-2d-matrix

// Testcases:
// case 0: [[1,3,5,7],[10,11,16,20],[23,30,34,60]]	3
// case 1: [[1,3,5,7],[10,11,16,20],[23,30,34,60]]	13

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m*n - 1
	for left <= right {
		mid := (right-left)/2 + left
		x := mid / n
		y := mid % n
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
