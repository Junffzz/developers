package main

import (
	"container/list"
	"fmt"
)

func main() {
	var root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	//ret := levelOrderOf(root)
	ret := dfsLevelOrder(root)
	fmt.Printf("levelOrderOf result: %v\n", ret)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
#102.Binary Tree
Level Order
二叉树层次遍历

解题思路：
1.BFS  O(N)
a.level=>queue
b.beatch process 代码

2.DFS  O(N)

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderOf(root *TreeNode) [][]int {
	return bfsLevelOrderOf(root)
}

func bfsLevelOrderOf(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	levels := make([][]int, 0)

	for len(queue) > 0 {
		n := len(queue)
		// 存放level层的数据
		level := make([]int, 0)
		for i := 0; i < n; i++ {
			root = queue[0]
			queue = queue[1:]

			level = append(level, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}

			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}

		levels = append(levels, level)
	}

	return levels
}

func dfsLevelOrder(root *TreeNode) [][]int {
	level := 0
	vals := make([][]int, 0)
	dfs(root, level, &vals)
	return vals
}

func dfs(root *TreeNode, level int, vals *[][]int) {
	// fmt.Println(root,level,vals)
	if root == nil {
		return
	}

	if len(*vals) <= level {
		*vals = append(*vals, []int{root.Val})
	} else {
		(*vals)[level] = append((*vals)[level], root.Val)
	}

	dfs(root.Left, level+1, vals)
	dfs(root.Right, level+1, vals)
}

// 另一个版本
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	visited := make(map[int]int)

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var curLevel []int
		count := queue.Len()
		for count > 0 {
			element := queue.Front()
			node := element.Value.(*TreeNode)

			if _, exist := visited[node.Val]; exist {
				continue
			}
			visited[node.Val]++

			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			queue.Remove(element)
			count--
		}
		result = append(result, curLevel)
	}
	return result
}
