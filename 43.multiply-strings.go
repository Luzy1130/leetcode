/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (31.40%)
 * Likes:    1398
 * Dislikes: 651
 * Total Accepted:    251.9K
 * Total Submissions: 775.7K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 * 
 * Example 1:
 * 
 * 
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 * 
 * Example 2:
 * 
 * 
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 * 
 * 
 * Note:
 * 
 * 
 * The length of both num1 and num2 is < 110.
 * Both num1 and num2 contain only digits 0-9.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 * 
 * 
 */

// @lc code=start
// package main
// import (
// 	"fmt"
// )

func reverseRune(runes []rune) {
	i := 0
	j := len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
}

func mult(num1 string, n rune, car int) string {
	var carry int = 0
	var res []rune
	// fmt.Println("aaa", num1, string(n), car)

	for i := len(num1)-1; i >= 0; i-- {
		num := int(n - '0') * int(rune(num1[i]) - '0') +  carry
		// fmt.Println(num)
		mod := num % 10
		res = append(res, rune(mod+int('0')))
		carry = num / 10
	}
	if carry > 0 {
		res = append(res, rune(carry) + '0')
	}
	reverseRune(res)

	for i := 0; i < car; i++ {
		res = append(res, '0')
	}
	return string(res)
}

func add(num1 string, num2 string) string {
	var res []rune
	var carry int = 0
	// fmt.Println("num1:",num1, ",num2:",num2)
	i := len(num1)-1
	j := len(num2)-1
	for ;i >= 0 && j >= 0; {
		num := int(rune(num1[i])-'0') + int(rune(num2[j])-'0') + carry
		mod := num % 10
		res = append(res, rune(mod)+'0')
		carry = num / 10
		i--
		j--
	}

	for ; i >= 0; i-- {
		num := int(rune(num1[i])-'0') + carry
		mod := num % 10
		res = append(res, rune(mod)+'0')
		carry = num / 10
	}
	for ; j >= 0; j-- {
		num := int(rune(num2[j])-'0') + carry
		mod := num % 10
		res = append(res, rune(mod)+'0')
		carry = num / 10
	}
	if carry > 0 {
		res = append(res, rune(carry) + '0')
	}
	reverseRune(res)

	return string(res)
}

// 大数乘法，不能将字符串转成数字
func multiply(num1 string, num2 string) string {
	// 其中一个数为0
    if num1 == "0" || num2 == "0" || len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	// num2的每一位与num1相乘并返回结果
	var tmps []string
	for i := len(num2)-1; i >= 0; i-- {
		tmp := mult(num1, rune(num2[i]), len(num2)-1-i)
		tmps = append(tmps, tmp)
	}

	// fmt.Println(tmps)
	// 将上面的每个结果相加，即是最后的乘积
	res := "0"
	for i := 0; i < len(tmps); i++ {
		res = add(res, tmps[i])
	}

	return res
}

// func main() {
// 	num1 := "123"
// 	num2 := "456"
// 	res := multiply(num1, num2)
// 	fmt.Println(res)
// }
// @lc code=end

