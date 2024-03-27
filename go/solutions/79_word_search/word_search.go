package wordsearch

// Title: Word Search

// Problem link: https://leetcode.com/problems/word-search

// Testcases:
// case 0: [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]	"ABCCED"
// case 1: [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]	"SEE"
// case 2: [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]	"ABCB"

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	xMove := [4]int{0, 0, 1, -1}
	yMove := [4]int{1, -1, 0, 0}
	var backTrack func(x, y int, word string) bool
	backTrack = func(x int, y int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if board[x][y] == word[0] {
			if len(word) == 1 {
				return true
			}
			for i := 0; i < 4; i++ {
				newX := x + xMove[i]
				newY := y + yMove[i]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && board[newX][newY] != '-' {
					temp := board[x][y]
					board[x][y] = '-'
					if backTrack(newX, newY, word[1:]) {
						return true
					}
					board[x][y] = temp
				}
			}
		}
		return false
	}
	for x, row := range board {
		for y, _ := range row {
			if backTrack(x, y, word) {
				return true
			}
		}
	}
	return false
}
