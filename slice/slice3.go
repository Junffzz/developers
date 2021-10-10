/*
@Time : 2020/11/13 18:54
@Author : ZhaoJunfeng
@File : slice3
*/
package main

import (
    "fmt"
)

func main() {
    // s2遍历
    s2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    fmt.Printf("len:%v\n", len(s2[:1]))
    fmt.Println("s2:", s2[3])

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0], nil, nil}
    var i = 0
    // 找中序遍历的根
    for i = 0; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
    root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

    return root
}

// 最大深度
func maxDepth(root *TreeNode) int {
    if root==nil {
        return 0
    }
    var level = 0
    level = dfs(root,level)
    return level
}

func dfs(root *TreeNode,level int) int{
    if root.Left !=nil {
        return dfs(root.Left,level+1)
    }

    if root.Right !=nil {
        return dfs(root.Right,level+1)
    }
    return level
}