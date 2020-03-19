/*
 * @lc app=leetcode id=137 lang=cpp
 *
 * [137] Single Number II
 *
 * https://leetcode.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (48.64%)
 * Likes:    1231
 * Dislikes: 307
 * Total Accepted:    198.5K
 * Total Submissions: 407.8K
 * Testcase Example:  '[2,2,3,2]'
 *
 * Given a non-empty array of integers, every element appears three times
 * except for one, which appears exactly once. Find that single one.
 * 
 * Note:
 * 
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 * 
 * Example 1:
 * 
 * 
 * Input: [2,2,3,2]
 * Output: 3
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [0,1,0,1,0,1,99]
 * Output: 99
 * 
 */

// @lc code=start
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int res = 0;
        std::vector<int> bits(32);
        for (int i = 0; i < 32; i++) {
            for (int num : nums) {
                // 将所有数的第i位相加
                bits[i] += (num >> i & 1);
                // 除3取余即是唯一数的第i位
                bits[i] %= 3;
            }
            // 将bit恢复到res中
            res |= (bits[i] << i);
        }
        return res;
    }
};
// @lc code=end

