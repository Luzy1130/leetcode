/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (40.27%)
 * Likes:    1153
 * Dislikes: 219
 * Total Accepted:    339.1K
 * Total Submissions: 835.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */
func addBinary(a string, b string) string {
	lena := len(a)
	lenb := len(b)

	var resLen int
	if lena > lenb {
		resLen = lena + 1
	} else {
		resLen = lenb + 1
	}

	resBytes := make([]byte, resLen)

	i := lena - 1
	j := lenb - 1
	k := resLen - 1
	var carry byte = 0
	for i >= 0 && j >= 0 {
		sumT := carry + (a[i] - '0') + (b[j] - '0')
		resBytes[k] = sumT%2 + '0'
		carry = sumT / 2
		i--
		j--
		k--
	}

	if i >= 0 {
		for i >= 0 {
			sumT := carry + (a[i] - '0')
			resBytes[k] = sumT%2 + '0'
			carry = sumT / 2
			i--
			k--
		}
	} else if j >= 0 {
		for j >= 0 {
			sumT := carry + (b[j] - '0')
			resBytes[k] = sumT%2 + '0'
			carry = sumT / 2
			j--
			k--
		}
	}

	if carry > 0 {
		resBytes[0] = carry + '0'
	} else {
		k = 1
	}

	return string(resBytes[k:])
}

