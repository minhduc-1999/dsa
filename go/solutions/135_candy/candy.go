package candy

// Title: Candy

// Problem link: https://leetcode.com/problems/candy

// Testcases:
// case 0: [1,0,2]
// case 1: [1,2,2]
// case 2: [1,2,1,0]

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	result := 0
	t1, t2, b := 0, 0, 0
	lastT2Height := 0
	i := 1
	for i < n {
		h1 := b - t1 + 1
		h2 := t2 - b + 1
		if ratings[i] == ratings[i-1] {
			// calc result
			result += h1 * (h1 + 1) / 2
			result += h2 * (h2 + 1) / 2
			result -= 1
			if lastT2Height > h1 {
				result -= h1
			} else {
				result -= lastT2Height
			}
			lastT2Height = 0
			t1, t2, b = i, i, i
		} else if ratings[i] < ratings[i-1] {
			if b != t2 {
				result += h1 * (h1 + 1) / 2
				result += h2 * (h2 + 1) / 2
				result -= 1
				if lastT2Height > (b - t1 + 1) {
					result -= b - t1 + 1
				} else {
					result -= lastT2Height
				}
				lastT2Height = t2 - b + 1
				t1, t2, b = i-1, i, i
			} else {
				b, t2 = i, i
			}

		} else {
			t2 = i
		}
		i += 1
	}
	result += (b - t1 + 1) * (b - t1 + 2) / 2
	result += (t2 - b + 1) * (t2 - b + 2) / 2
	result -= 1
	if lastT2Height > (b - t1 + 1) {
		result -= b - t1 + 1
	} else {
		result -= lastT2Height
	}
	return result
}
