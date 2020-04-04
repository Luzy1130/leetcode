/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (44.11%)
 * Likes:    2489
 * Dislikes: 44
 * Total Accepted:    251.2K
 * Total Submissions: 564.5K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Implement a trie with insert, search, and startsWith methods.
 * 
 * Example:
 * 
 * 
 * Trie trie = new Trie();
 * 
 * trie.insert("apple");
 * trie.search("apple");   // returns true
 * trie.search("app");     // returns false
 * trie.startsWith("app"); // returns true
 * trie.insert("app");   
 * trie.search("app");     // returns true
 * 
 * 
 * Note:
 * 
 * 
 * You may assume that all inputs are consist of lowercase letters a-z.
 * All inputs are guaranteed to be non-empty strings.
 * 
 * 
 */

// 前缀树/字典树

// @lc code=start
type TrieNode struct {
	isKey bool
	children []*TrieNode
}

type Trie struct {
    root *TrieNode
}

func newTrieNode() *TrieNode{
	node := &TrieNode{
		isKey: false,
		children: make([]*TrieNode, 26),
	}
	return node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{root: newTrieNode()}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this.root
	var nextNode *TrieNode
	for i, _ := range word {
		index := int(word[i]-'a')
		if node.children[index] == nil {
			nextNode = newTrieNode()
			node.children[index] = nextNode
		} else {
			nextNode = node.children[index]
		}

		node = nextNode
	}
	if nextNode != nil {
		nextNode.isKey = true
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for i, _ := range word {
		index := int(word[i] - 'a')
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return node.isKey
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this.root
	for i, _ := range prefix {
		index := int(prefix[i] - 'a')
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

