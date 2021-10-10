package main

import (
    "math"
)

func main() {

}

/**
#104 题目：二叉树的最大深度和最小深度
示例：
给定二叉树[3,9,20,null,null,15,7,null,4],
返回它的最大深度4，最小深度2

解法：
1.递归（不成体系，排除）
2.DFS O(N) N为节点数
3.BFS O(N) 较好
*/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 最大深度
func maxDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return int(float64(1) + math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

func minDepth(root *TreeNode) int {
    if root==nil {
        return 0
    }
    //todo
    return 0
}