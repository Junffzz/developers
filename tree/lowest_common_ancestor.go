package main

func main() {

}

//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
//}

/**
二叉树/二叉搜索树的最近公共祖先
实例1：
输入：root = [6,2,8,0,4,7,9,null,null,3,5], p=2,q=8
输出：6

实例2：
输入：root=[6,2,8,0,4,7,9,null,null,3,5], p=2,q=4
输出：2

解法：
1.Path：O(N),path1,path2=>LCA
2.Recursion: 辅助函数_findPorQ(root,p,q)
*/
// 二叉树
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    //说明分叉了，公共祖先为root
    if left != nil && right != nil {
        return root
    } else if left != nil {
        return left
    } else {
        return right
    }
    //return root
}

// 二叉搜索树（递归或非递归）
func lowestCommonAncestorOfBinarySearchTree(root, p, q *TreeNode) *TreeNode {
    for {
        if  root.Val>p.Val && root.Val > q.Val {
            root = root.Left
        } else if root.Val<p.Val && root.Val<q.Val {
            root = root.Right
        } else {
            return root
        }
    }
}
