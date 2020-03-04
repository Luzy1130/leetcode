/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (26.94%)
 * Likes:    4657
 * Dislikes: 203
 * Total Accepted:    441K
 * Total Submissions: 1.5M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design and implement a data structure for Least Recently Used (LRU) cache.
 * It should support the following operations: get and put.
 * 
 * get(key) - Get the value (will always be positive) of the key if the key
 * exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present.
 * When the cache reached its capacity, it should invalidate the least recently
 * used item before inserting a new item.
 * 
 * The cache is initialized with a positive capacity.
 * 
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 * 
 * Example:
 * 
 * 
 * LRUCache cache = new LRUCache( 2  );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // returns 1
 * cache.put(3, 3);    // evicts key 2
 * cache.get(2);       // returns -1 (not found)
 * cache.put(4, 4);    // evicts key 1
 * cache.get(1);       // returns -1 (not found)
 * cache.get(3);       // returns 3
 * cache.get(4);       // returns 4
 * 
 * 
 * 
 * 
 */

// @lc code=start
// package main
// import "fmt"

type Node struct {
	Key   int
	Value int
	Pre	 *Node
	Next *Node
}

type LRUCache struct {
	Capacity int
	Len		 int
	Index	 map[int]*Node
	ListHead *Node	// 使用一个哨兵节点
	ListTail *Node	//指向尾结点，尾结点也是LRU节点
}


func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.Capacity = capacity
	cache.Len = 0
	cache.ListHead = &Node{}
	cache.ListTail = nil
	cache.Index = make(map[int]*Node)
	return cache
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.Index[key]; ok {
		// fmt.Println("get", key)
		// 访问后将元素作为表头
		node := this.Index[key]
		if node.Pre != this.ListHead {
			// 将最近访问的节点移到表头
			if node == this.ListTail {
				this.ListTail = node.Pre
			}

			node.Pre.Next = node.Next
			if node.Next != nil {
				node.Next.Pre = node.Pre
			}
			node.Pre = this.ListHead

			this.ListHead.Next.Pre = node
			node.Next = this.ListHead.Next
			this.ListHead.Next = node
			// fmt.Println("tail", this.ListTail.Key)
		}

		return node.Value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if this.Get(key) != -1 {
		// Get已经更新了访问时间
		this.Index[key].Value = value
		return
	}
	newNode := &Node{}
	newNode.Value = value
	newNode.Key = key
	newNode.Pre = this.ListHead
	newNode.Next = this.ListHead.Next
	if this.ListHead.Next != nil {
		this.ListHead.Next.Pre = newNode
	}
	this.ListHead.Next = newNode

	if this.ListTail == nil {
		this.ListTail = newNode
	}

	// fmt.Println("put", key)

	// 驱逐LRU元素
	if this.Len == this.Capacity {
		oldKey := this.ListTail.Key
		this.ListTail.Pre.Next = this.ListTail.Next
		this.ListTail = this.ListTail.Pre
		delete(this.Index, oldKey)
		// fmt.Println("delete: ", oldKey)
	} else {
		this.Len++
	}

	this.Index[key] = newNode
}

// func main() {
// 	obj := Constructor(3)

// 	obj.Put(1,1)
// 	obj.Put(2,2)
// 	obj.Put(3,3)
// 	obj.Put(4,4)

// 	fmt.Println(obj.Get(4))
// 	fmt.Println(obj.Get(3))
// 	fmt.Println(obj.Get(2))
// 	fmt.Println(obj.Get(1))

// 	obj.Put(5, 5)	// 4驱逐
	
// 	fmt.Println(obj.Get(1))
// 	fmt.Println(obj.Get(2))
// 	fmt.Println(obj.Get(3))
// 	fmt.Println(obj.Get(4))
// 	fmt.Println(obj.Get(5))
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

