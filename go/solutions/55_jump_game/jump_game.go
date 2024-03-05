package jumpgame

// Title: Jump Game

// Problem link: https://leetcode.com/problems/jump-game

// Testcases:
// case 0: [2,3,1,1,4]
// case 1: [3,2,1,0,4]

func canJump(nums []int) bool {
  maxIdx := 0
  for i, distance := range nums {
    if i > maxIdx {
      return false
    }
    maxIdx = max(maxIdx, distance + i)
  }
  return true
}
