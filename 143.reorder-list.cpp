/*
 * @lc app=leetcode id=143 lang=cpp
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (34.56%)
 * Likes:    1474
 * Dislikes: 99
 * Total Accepted:    203.8K
 * Total Submissions: 586.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a singly linked list L: L0→L1→…→Ln-1→Ln,
 * reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
 * 
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 * 
 * Example 1:
 * 
 * 
 * Given 1->2->3->4, reorder it to 1->4->2->3.
 * 
 * Example 2:
 * 
 * 
 * Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
 * 
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
class Solution {
public:
    // // 使用递归来实现, 效率偏低
    // void reorderListRecursion(ListNode *head) {
    //     if (head == nullptr || head->next == nullptr || head->next->next == nullptr) {
    //         return;
    //     }

    //     ListNode *pos1 = head;
    //     ListNode *pos2 = head->next;
    //     ListNode *pre = head;
    //     while (pos2->next != nullptr) {
    //         pre = pos2;
    //         pos2 = pos2->next;
    //     }

    //     ListNode *nextHead = pos1->next;
    //     pos1->next = pos2;
    //     pos2->next = nextHead;
    //     pre->next = nullptr;

    //     reorderListRecursion(nextHead);
    // }

    void reorderListNonRecursion(ListNode *head) {
        if (head == nullptr || head->next == nullptr || head->next->next == nullptr) {
            return;
        }

        ListNode* fast = head;
        ListNode* slow = head;
        while(fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
        }

        // slow后面的节点需要向前插入
        ListNode* behind = slow->next;
        slow->next = nullptr;

        // 将behind倒序
        ListNode *pre = nullptr;
        ListNode *cur = behind;
        while(cur->next != nullptr) {
            ListNode* next = cur->next;
            cur->next = pre;
            pre = cur;
            cur = next;
        }
        cur->next = pre;

        // 将head和cur两个链表合并即可
        ListNode *pos1 = head;
        ListNode *pos2 = cur;
        while (pos1 != nullptr && pos2 != nullptr) {
            ListNode* next1 = pos1->next;
            ListNode* next2 = pos2->next;

            pos1->next = pos2;
            pos2->next = next1;
            pos1 = next1;
            pos2 = next2;
        }
    }

    void reorderList(ListNode* head) {
        // reorderListRecursion(head);

        reorderListNonRecursion(head);
    }
};
// @lc code=end

