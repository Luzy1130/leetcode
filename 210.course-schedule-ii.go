/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 *
 * https://leetcode.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (35.98%)
 * Likes:    1612
 * Dislikes: 107
 * Total Accepted:    214.7K
 * Total Submissions: 563.3K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 * 
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 * 
 * Given the total number of courses and a list of prerequisite pairs, return
 * the ordering of courses you should take to finish all courses.
 * 
 * There may be multiple correct orders, you just need to return one of them.
 * If it is impossible to finish all courses, return an empty array.
 * 
 * Example 1:
 * 
 * 
 * Input: 2, [[1,0]] 
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you
 * should have finished   
 * course 0. So the correct course order is [0,1] .
 * 
 * Example 2:
 * 
 * 
 * Input: 4, [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,1,2,3] or [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you
 * should have finished both     
 * ⁠            courses 1 and 2. Both courses 1 and 2 should be taken after you
 * finished course 0. 
 * So one correct course order is [0,1,2,3]. Another correct ordering is
 * [0,2,1,3] .
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
}

func BuildGraph(num int, prerequisites [][]int) *Graph {
	g := &Graph{}
	g.Num = num
	g.LinkedList = make(map[int][]int)
	g.Nodes = make(map[int]bool)

	for i:=0; i < num; i++ {
		g.Nodes[i] = true
	}

	// fmt.Println(num, prerequisites)
	for _, prerequisite := range prerequisites {
		s := prerequisite[1]
		t := prerequisite[0]

		g.AddEdge(s, t)
	}
	return g
}
func findOrder(numCourses int, prerequisites [][]int) []int {
	defaultRes := []int{}

	for i := numCourses-1; i >= 0; i-- {
		defaultRes = append(defaultRes, i)
	}
	if len(prerequisites) == 0 {
		return defaultRes
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
		return []int{}
	}

	res := []int{}
	visited := make(map[int]bool)
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]
		res = append(res, n)
		if _, ok := visited[n]; ok {
			return []int{}
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
			return []int{}
		}
	}

	return res
}
// @lc code=end

