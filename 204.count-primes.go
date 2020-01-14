/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 *
 * https://leetcode.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (29.69%)
 * Likes:    1471
 * Dislikes: 492
 * Total Accepted:    294.7K
 * Total Submissions: 970.8K
 * Testcase Example:  '10'
 *
 * Count the number of prime numbers less than a non-negative number, n.
 * 
 * Example:
 * 
 * 
 * Input: 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 * 
 * 
 */

// @lc code=start
func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func countPrimes(n int) int {
	notPrimeMap := make(map[int]bool)
	cnt := 0
	for i := 2; i < n; i++ {
		if _, ok := notPrimeMap[i]; !ok {
			cnt++
			for j := i*i; j <n ; j+=i {
				notPrimeMap[j] = true
			}
		}
	}

	return cnt
}
// @lc code=end

