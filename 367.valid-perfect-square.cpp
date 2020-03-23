/*
 * @lc app=leetcode id=367 lang=cpp
 *
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (40.91%)
 * Likes:    627
 * Dislikes: 137
 * Total Accepted:    147K
 * Total Submissions: 359.3K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 * 
 * Note: Do not use any built-in library function such as sqrt.
 * 
 * Example 1:
 * 
 * 
 * 
 * Input: 16
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 14
 * Output: false
 * 
 * 
 * 
 */

// @lc code=start
// #include <iostream>

class Solution {
public:
    bool isPerfectSquare(int num) {
        int low = 1;
        int high = num;
        if (high == INT_MAX) {
            high -= 1;
        }

        while (low <= high) {
            int mid = (low + high) / 2;
            int rem = num / mid;
            if (rem == mid && rem * mid == num) {
                return true;
            }
            if (rem < mid) {
                high = mid-1;
            } else {
                low = mid+1;
            }
        }
        return false;
    }
};

// int main(void) {
//     Solution s;
//     std::cout << "1: " << s.isPerfectSquare(1) << std::endl;
//     std::cout << "2: " << s.isPerfectSquare(2) << std::endl;
//     std::cout << "3: " << s.isPerfectSquare(3) << std::endl;
//     std::cout << "4: " << s.isPerfectSquare(4) << std::endl;
//     std::cout << "5: " << s.isPerfectSquare(5) << std::endl;
//     std::cout << "6: " << s.isPerfectSquare(6) << std::endl;
//     std::cout << "7: " << s.isPerfectSquare(7) << std::endl;
//     std::cout << "8: " << s.isPerfectSquare(8) << std::endl;
//     std::cout << "9: " << s.isPerfectSquare(9) << std::endl;
//     std::cout << "10: " << s.isPerfectSquare(10) << std::endl;
//     std::cout << "11: " << s.isPerfectSquare(11) << std::endl;
//     return 0;
// }
// @lc code=end

