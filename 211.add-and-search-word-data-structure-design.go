/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Add and Search Word - Data structure design
 *
 * https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
 *
 * algorithms
 * Medium (34.25%)
 * Likes:    1409
 * Dislikes: 75
 * Total Accepted:    159.3K
 * Total Submissions: 460.8K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * Design a data structure that supports the following two operations:
 * 
 * 
 * void addWord(word)
 * bool search(word)
 * 
 * 
 * search(word) can search a literal word or a regular expression string
 * containing only letters a-z or .. A . means it can represent any one
 * letter.
 * 
 * Example:
 * 
 * 
 * addWord("bad")
 * addWord("dad")
 * addWord("mad")
 * search("pad") -> false
 * search("bad") -> true
 * search(".ad") -> true
 * search("b..") -> true
 * 
 * 
 * Note:
 * You may assume that all words are consist of lowercase letters a-z.
 * 
 */

// 前缀树/字典树
package main
import "fmt"

// @lc code=start
type Node struct {
	isKey bool
	children []*Node
}

func newNode() *Node {
	node := &Node{
		isKey: false,
		children: make([]*Node, 26),
	}
	return node
}

type WordDictionary struct {
    root *Node
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	w := WordDictionary{}
	w.root = newNode()
	return w
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
    node := this.root
	var nextNode *Node
	for i, _ := range word {
		index := int(word[i]-'a')
		if node.children[index] == nil {
			nextNode = newNode()
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

func (n *Node)SearchImpl(word string) bool {
	// fmt.Println(word)
	if len(word) == 0 {
		return n.isKey
	}
	wordTmp := []rune(word)

	node := n
	for i, _ := range word {
		if word[i] == '.' {
			for _, child := range node.children {
				if child != nil {
					res := child.SearchImpl(string(wordTmp[i+1:]))
					if res == true {
						return true
					}
				}
			}
			// fmt.Println("cant find")
			return false
		} else {
			index := int(word[i] - 'a')
			if node.children[index] == nil {
				return false
			}
			// fmt.Println(wordTmp[i])
			node = node.children[index]
		}
	}

	return node.isKey
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this.root

	return node.SearchImpl(word)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

func main() {
	w := Constructor()

	w.AddWord("bad")
	w.AddWord("dad")
	w.AddWord("mad")

	res := w.Search("b..")
	fmt.Println(res)
}