/*
 * @lc app=leetcode id=350 lang=cpp
 *
 * [350] Intersection of Two Arrays II
 *
 * https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (50.38%)
 * Likes:    1090
 * Dislikes: 354
 * Total Accepted:    298.9K
 * Total Submissions: 593K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 * 
 * Example 1:
 * 
 * 
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2,2]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [4,9]
 * 
 * 
 * Note:
 * 
 * 
 * Each element in the result should appear as many times as it shows in both
 * arrays.
 * The result can be in any order.
 * 
 * 
 * Follow up:
 * 
 * 
 * What if the given array is already sorted? How would you optimize your
 * algorithm?
 * What if nums1's size is small compared to nums2's size? Which algorithm is
 * better?
 * What if elements of nums2 are stored on disk, and the memory is limited such
 * that you cannot load all elements into the memory at once?
 * 
 * 
 */

// @lc code=start
// 1. 如果数组已经排序，是不是可以优化；
// 2. nums1和num2的长度是否会影响算法的性能；
// 3. 如果数组存在磁盘且很大，而内存不够大，怎么优化这个算法
class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        std::vector<int> res;
        if (nums1.empty() || nums2.empty()) {
            return res;
        }
        // 先排序, 可以考虑适合外部排序的算法
        std::sort(nums1.begin(), nums1.end());
        std::sort(nums2.begin(), nums2.end());

        std::unordered_map<int, int> numMap;
        int i = 0;
        int j = 0;
        for (; i < nums1.size(); i++) {
            if (nums1[i] != nums1[j]) {
                numMap[nums1[j]] = (i - j);
                j = i;
            }
        }
        numMap[nums1[j]] = i - j;

        for (i = 0; i < nums2.size(); i++) {
            int num = nums2[i];
            if (numMap.find(num) != numMap.end() && numMap[num] != 0) {
                res.push_back(num);
                numMap[num]--;
            }
        }
        return res;
    }
};
// @lc code=end

