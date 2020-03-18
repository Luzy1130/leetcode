/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;
    
    Node() {
        val = 0;
        neighbors = vector<Node*>();
    }
    
    Node(int _val) {
        val = _val;
        neighbors = vector<Node*>();
    }
    
    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};
*/
class Solution {
public:
    // 应该是考察图的遍历
    Node* cloneGraph(Node* node) {
        // 用bfs遍历吧
        std::queue<Node*> q;
        std::queue<Node*> nq;
        std::map<Node*, Node*> hash;
        
        if (node == nullptr) {
            return nullptr;
        }
        
        Node* res = new Node(node->val);
        q.push(node);
        nq.push(res);
        //std::cout << (*node).val << std::endl;
        //std::cout << (*res).val << std::endl;
        hash[node] = res;
        while(!q.empty()) {
            std::queue<Node*> tmpQ;
            std::queue<Node*> tmpNQ;
            while(!q.empty()) {
                Node* tmpN = q.front();
                Node* tmpR = nq.front();
                
                for(int i = 0; i < tmpN->neighbors.size(); i++) {
                    Node* nb = tmpN->neighbors[i];
                    Node* newNB = nullptr;
                    if (hash.find(nb) == hash.end()) {
                        newNB = new Node(nb->val);
                        //std::cout << (*nb).val << std::endl;
                        //std::cout << (*newNB).val << std::endl;
                        tmpQ.push(nb);
                        tmpNQ.push(newNB);
                        hash[nb] = newNB;
                    } else {
                        newNB = hash[nb];
                    }
                    tmpR->neighbors.push_back(newNB);
                }
                
                q.pop();
                nq.pop();
            }
            q = tmpQ;
            nq = tmpNQ;
        }
        return res;
    }
};