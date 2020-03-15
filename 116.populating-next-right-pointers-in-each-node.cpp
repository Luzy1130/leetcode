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
        std::queue<Node *> q;
        
        if (root != nullptr) {
            q.push(root);
        }
        
        while(!q.empty()) {
            std::queue<Node *> tmp;
            Node* node = q.front();
            while(!q.empty()) {
                q.pop();
                if (!q.empty()) {
                    node->next = q.front();
                }
                if (node->left != nullptr) {
                    tmp.push(node->left);
                }
                if (node->right != nullptr) {
                    tmp.push(node->right);
                }
                node = node->next;
            }
            q = tmp;
        }
        
        return root;
    }
};
