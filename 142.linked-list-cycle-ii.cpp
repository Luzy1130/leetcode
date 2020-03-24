/*
 * @lc app=leetcode id=142 lang=cpp
 *
 * [142] Linked List Cycle II
 *
 * https://leetcode.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (35.44%)
 * Likes:    2189
 * Dislikes: 181
 * Total Accepted:    284.1K
 * Total Submissions: 799.2K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, return the node where the cycle begins. If there is no
 * cycle, return null.
 * 
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 * 
 * Note: Do not modify the linked list.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 * 
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 * 
 * 
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 * 
 * 
 * 
 * 
 * 
 * 
 * Follow-up:
 * Can you solve it without using extra space?
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

// Suppose,
// Length of cycle, C = 8
// Length of non cycle part of list, M = 3
// When slow enters first node of cycle, dist bw slow and fast is D = 3.
// Now, fast has to catch up slow with 2x speed. I.E relative to slow, fast has to travel C-D = 5, nodes. Since, fast is 2x of slow and relative to slow it travels 5 nodes, so originally it will need to travel 2*(C-D) nodes and slow need to travel (C-D) nodes for them to meet.

// image
// It is easy to point out that D is equal to Length of non cycle part of list, M.

// image

// At meeting point C-D (slow has to travel C-D nodes to meet), slow is D nodes away from cycle starting.
// Since D = M, slow is M nodes away from cycly starting point. So is the head pointer. If they both move one node at a time, they will meet at cycle starting point. 
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        ListNode *fast = head;
        ListNode *slow = head;

        while (fast != nullptr && slow != nullptr) {
            fast = fast->next;
            slow = slow->next;
            if (fast != nullptr) {
                fast = fast->next;
            } else {
                break;
            }

            if (fast == slow) {
                break;
            }
        }
        if (fast == nullptr) {
            return nullptr;
        }

        ListNode *pos = head;
        while (pos != slow) {
            pos = pos->next;
            slow = slow->next;
        }
        return pos;
    }
};
// @lc code=end

