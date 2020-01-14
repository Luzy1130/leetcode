/*
 * @lc app=leetcode id=278 lang=cpp
 *
 * [278] First Bad Version
 *
 * https://leetcode.com/problems/first-bad-version/description/
 *
 * algorithms
 * Easy (31.01%)
 * Likes:    897
 * Dislikes: 519
 * Total Accepted:    282.2K
 * Total Submissions: 870.3K
 * Testcase Example:  '5\n4'
 *
 * You are a product manager and currently leading a team to develop a new
 * product. Unfortunately, the latest version of your product fails the quality
 * check. Since each version is developed based on the previous version, all
 * the versions after a bad version are also bad.
 * 
 * Suppose you have n versions [1, 2, ..., n] and you want to find out the
 * first bad one, which causes all the following ones to be bad.
 * 
 * You are given an API bool isBadVersion(version) which will return whether
 * version is bad. Implement a function to find the first bad version. You
 * should minimize the number of calls to the API.
 * 
 * Example:
 * 
 * 
 * Given n = 5, and version = 4 is the first bad version.
 * 
 * call isBadVersion(3) -> false
 * call isBadVersion(5) -> true
 * call isBadVersion(4) -> true
 * 
 * Then 4 is the first bad version. 
 * 
 * 
 */
// 1,2,3,4,5,6
// @lc code=start
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int i = 0;
        int j = n;
        // 这里int溢出是个异常点，需要注意
        int checkIndex = (int)(((double)i+(double)j)/2);

        while (j-i > 1) {
            if (isBadVersion(checkIndex)) {
                j = checkIndex;
            } else {
                i = checkIndex;
            }
            checkIndex = (int)(((double)i+(double)j)/2);
        }
        return j;
    }
};
// @lc code=end

