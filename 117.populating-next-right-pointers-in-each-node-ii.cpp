/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/
class Solution {
public:
    Node* connect(Node* root) {
        if (root == nullptr) {
            return root;
        }

        Node* pNode = root;
        Node* pLeftNode = nullptr;
        Node* pMostLeftNode = nullptr;

        while(pNode != nullptr) {
            if(pNode->left != nullptr) {
                if (pMostLeftNode == nullptr) {
                    pMostLeftNode = pNode->left;
                }

                if (pLeftNode != nullptr) {
                    pLeftNode->next = pNode->left;
                }
                pLeftNode = pNode->left;
            }
            if (pNode->right != nullptr) {
                if (pMostLeftNode == nullptr) {
                    pMostLeftNode = pNode->right;
                }
                if (pLeftNode != nullptr) {
                    pLeftNode->next = pNode->right;
                }
                pLeftNode = pNode->right;
            }

            if (pNode->next != nullptr) {
                pNode = pNode->next;
            } else {
                // 换层
                pLeftNode = nullptr;
                pNode = pMostLeftNode;
                pMostLeftNode = nullptr;
            }
        }

        return root;
    }
};