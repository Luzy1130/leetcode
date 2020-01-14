/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */


func reverse(x int) int {
	var result int = 0

	for ; x != 0; x = x/10 {
		tmp := x % 10
		if tmp > 0 {
			if (math.MaxInt32 - tmp) / 10 < result {
				return 0
			}
		} else {
			if (math.MinInt32 - tmp) / 10 > result {
				return 0
			}
		}
		result = result * 10 + x % 10
	}

	return result
}

