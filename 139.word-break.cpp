/*
 * @lc app=leetcode id=139 lang=cpp
 *
 * [139] Word Break
 *
 * https://leetcode.com/problems/word-break/description/
 *
 * algorithms
 * Medium (38.39%)
 * Likes:    3525
 * Dislikes: 193
 * Total Accepted:    475.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of
 * non-empty words, determine if s can be segmented into a space-separated
 * sequence of one or more dictionary words.
 * 
 * Note:
 * 
 * 
 * The same word in the dictionary may be reused multiple times in the
 * segmentation.
 * You may assume the dictionary does not contain duplicate words.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 * 
 * 
 */

// @lc code=start
// #include <vector>
// #include <iostream>
// using namespace std;

class Solution {
public:
    // // dfs会导致冗余计算过多
    // bool wordBreakDfsImpl(string s, unordered_map<string, bool> &wordDict) {
    //     if (wordDict.find(s) != wordDict.end()) {
    //         return true;
    //     }

    //     for (int i = 0; i < s.length(); i++) {
    //         string tmp(s.begin(), s.begin()+i+1);
    //         // cout << "preffix: " << tmp << endl;
    //         if (wordDict.find(tmp) != wordDict.end()) {
    //             string suffix(s.begin()+i+1, s.end());
    //             // cout << "suffix: " << suffix << endl;
    //             // 这一个递归会产生很多冗余重复的计算
    //             bool tmpRes = wordBreakDfsImpl(string(s.begin()+i+1, s.end()), wordDict);
    //             if (tmpRes) {
    //                 return true;
    //             }
    //         }
    //     }
    //     return false;
    // }

    // // dp已经记录过的字符串
    // bool wordBreakDpImpl1(string s, unordered_map<string, bool> &wordDict, int start, int end) {
    //     cout << "start: " << start << " end: " << end << endl; 
    //     if (start > end) {
    //         return false;
    //     }
    //     if (dp[start][end] == 1) {
    //         return true;
    //     } else if (dp[start][end] == -1) {
    //         return false;
    //     }

    //     for (int i = 0; i < s.length(); i++) {
    //         string tmp(s.begin(), s.begin()+i+1);
    //         if (dp[start][start+i] == 1 || wordDict.find(tmp) != wordDict.end()) {
    //             dp[start][start+i] = 1;
    //             if (i == s.length()-1) {
    //                 return true;
    //             }
    //             string suffix(s.begin()+i+1, s.end());
    //             bool tmpRes = wordBreakDpImpl1(string(s.begin()+i+1, s.end()), wordDict, start+i+1, end);
    //             if (tmpRes) {
    //                 dp[start][end] = 1;
    //                 return true;
    //             }
    //         } else {
    //             dp[start][start+i] = -1;
    //         }
    //     }

    //     dp[start][end] = -1;
    //     return false;
    // }

    // 使用dict去匹配会更快
    bool wordBreakDpImpl2(string s, vector<string> &wordDict) {
        if (s.empty() || wordDict.empty()) {
            return false;
        }

        vector<bool> dp(s.size()+1, false);
        dp[0] = true;

        for (int i = 0; i < s.length(); i++) {
            for (string word : wordDict) {
                int len = word.length();
                if ((i+1) >= len && dp[i+1-len] == true && s.substr(i+1-len, len) == word) {
                    // cout << word << endl;
                    dp[i+1] = true;
                    break;
                }
            }
        }
        return dp[s.length()];
    }
    
    bool wordBreak(string s, vector<string>& wordDict) {
        // unordered_map<string, bool> wordMap;

        // for (string word : wordDict) {
        //     wordMap[word] = true;
        // }

        // // 1. method1: dfg
        // return wordBreakDfsImpl(s, wordMap);

        // // 2. dp1
        // int len = s.length();
        // dp = new int*[len];
        // for (int i = 0; i < len; i++) {
        //     dp[i] = new int[len];
        //     for (int j = 0; j < len; j++) {
        //         dp[i][j] = 0;
        //     }
        // }
        // return wordBreakDpImpl1(s, wordMap, 0, len-1);

        // 2. dp2
        return wordBreakDpImpl2(s, wordDict);
    }
private:
    // int** dp;   // dp[i][j]记录是否满足要求
};


// int main() {
//     Solution s;
//     vector<string> dic;
//     dic.push_back("apple");
//     dic.push_back("pen");
//     s.wordBreak("applepenapple", dic);
// }
// @lc code=end

