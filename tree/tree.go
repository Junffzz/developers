package main

import "math"

func main() {

}

/**
#98.验证二叉搜索树
示例1：
输入：[3,1,5,null,null,2]
示例2：
输入：[5,1,4,null,null,3,6]
输出：false

解法：
1.中序遍历：In-order=>array 升序    O(N)
2.递归 Recursion：validate(...,min,max){     O(N)
    max<-validate(node,left) 左子树最大值
    min<-validate(node,right) 右子树最小值

    max<root,min>root
}
*/
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, left, right int) bool {
    if root == nil {
        return true
    }

    if left >= root.Val || right <= root.Val {
        return false
    }

    return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}
