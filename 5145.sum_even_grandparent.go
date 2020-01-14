/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var levelNodes [][]*TreeNode
	
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelNodes = append(levelNodes, queue)
		var tmpQ []*TreeNode

		for _, n := range queue {
			if n.Left != nil {
				tmpQ = append(tmpQ, n.Left)
			}
			if n.Right != nil {
				tmpQ = append(tmpQ, n.Right)
			}
		}
		queue = tmpQ
	}

	sum := 0
	for i := 0; i <len(levelNodes)-2; i++ {
		tmpNodes := levelNodes[i]
		for _, n := range tmpNodes {
			if n.Val % 2 == 0 {
				lc := n.Left
				rc := n.Right

				if lc != nil && lc.Left != nil {
					sum += lc.Left.Val
				}
				if lc != nil && lc.Right != nil {
					sum += lc.Right.Val
				}
				if rc != nil && rc.Left != nil {
					sum += rc.Left.Val
				}
				if rc != nil && rc.Right != nil {
					sum += rc.Right.Val
				}
			}
		}
	}
	return sum
}

func main() {
	root := &TreeNode{6, nil, nil}
	res := sumEvenGrandparent(root)
	fmt.Println(res)
}