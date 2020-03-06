/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (38.99%)
 * Likes:    2900
 * Dislikes: 146
 * Total Accepted:    326.8K
 * Total Submissions: 799.8K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 * 
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 * 
 * Given the total number of courses and a list of prerequisite pairs, is it
 * possible for you to finish all courses?
 * 
 * Example 1:
 * 
 * 
 * Input: 2, [[1,0]] 
 * Output: true
 * Explanation: There are a total of 2 courses to take. 
 * To take course 1 you should have finished course 0. So it is possible.
 * 
 * Example 2:
 * 
 * 
 * Input: 2, [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take. 
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should
 * also have finished course 1. So it is impossible.
 * 
 * 
 * Note:
 * 
 * 
 * The input prerequisites is a graph represented by a list of edges, not
 * adjacency matrices. Read more about how a graph is represented.
 * You may assume that there are no duplicate edges in the input
 * prerequisites.
 * 
 * 
 */

// @lc code=start
// package main

// import "fmt"

// 通过拓扑排序，找出是否有环,有环则返回false

type Graph struct {
	Num int
	LinkedList map[int][]int
	Nodes map[int]bool
}

func (g *Graph) AddEdge(s, t int) {
	if _, ok := g.LinkedList[s]; !ok {
		g.LinkedList[s] = []int{}
	}
	g.LinkedList[s] = append(g.LinkedList[s], t)

	g.Nodes[s] = true
	g.Nodes[t] = true
}

func BuildGraph(num int, prerequisites [][]int) *Graph {
	g := &Graph{}
	g.Num = num
	g.LinkedList = make(map[int][]int)
	g.Nodes = make(map[int]bool)

	// fmt.Println(num, prerequisites)
	for _, prerequisite := range prerequisites {
		s := prerequisite[1]
		t := prerequisite[0]

		g.AddEdge(s, t)
	}
	return g
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	// fmt.Println("---begin----")
	g := BuildGraph(numCourses, prerequisites)
	
	// fmt.Println(g)
	// 统计每个节点的入度
	inDegree := make(map[int]int)
	for _, ts := range g.LinkedList {
		for _, t := range ts {
			inDegree[t]++
		}
	}

	queue := []int{}
	for n, _ := range g.Nodes {
		if v, ok := inDegree[n]; !ok || v == 0 {
			queue = append(queue, n)
		}
	}

	if len(queue) == 0 {
		return false
	}

	visited := make(map[int]bool)
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = true
		for _, v := range g.LinkedList[n] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
		// fmt.Println(queue)
	}

	for v, _ := range g.Nodes {
		if _, ok := visited[v]; !ok {
			return false
		}
	}

	return true
}

// func main() {
// 	res := canFinish(2, [][]int{[]int{1, 0}})
// 	fmt.Println(res)
// }
// @lc code=end

