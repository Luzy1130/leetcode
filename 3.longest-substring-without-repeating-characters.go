/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.77%)
 * Likes:    7230
 * Dislikes: 425
 * Total Accepted:    1.2M
 * Total Submissions: 4.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "abcabcbb"
 * Output: 3 
 * Explanation: The answer is "abc", with the length of 3. 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3. 
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
func method1(s string) int {
	res := 0
	// lSubPos := 0 // 用于记录最长不重复子串的起始位置
	subPos := 0	 // 用于滑动的序列窗口
	charMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if v, ok := charMap[s[i]]; ok {
			// 当前滑动窗口内如果有重复的，则将窗口滑动到根据重复字符的位置左边
			subPos += (v + 1)
			// 更新charMap中每个字符的位置
			charMap = make(map[byte]int)
			for pos := subPos; pos <= i; pos++ {
				charMap[s[pos]] = pos - subPos
			}
		} else {
			// 在当前滑动窗口内没有重复字符的情况
			dis := i - subPos
			if dis+1 > res {
				// 如果找到了更长的子串，更新其实位置和长度
				res = dis + 1
				// lSubPos = subPos
			}
			charMap[s[i]] = dis
		}
	}
	// fmt.Println(s[lSubPos:(lSubPos+res)])
	return res
}

func method2(s string) int {
	charMap := make(map[byte]int)	// byte->pos
	start := 0	// 标记滑动窗口的其实位置
	maxLen := 0

	for i := 0; i < len(s); i++ {
		v, ok := charMap[s[i]];
		if (ok && v+1 > start) {
			start = v + 1
		}
		charMap[s[i]] = i
		if maxLen < i - start + 1 {
			maxLen = i - start + 1
		}
	}

	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	// return methods1(s)
	return method2(s)
}
// @lc code=end

