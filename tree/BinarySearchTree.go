package main

import (
    "fmt"
    "math"
    "reflect"
)

func main() {
    var tree = &treeNode{val: 33}
    tree.insert(16).
        insert(50).
        insert(13).
        insert(18).
        insert(34).
        insert(58).
        insert(15).
        insert(17).
        insert(25).
        insert(51).
        insert(66).
        insert(19).
        insert(27).
        insert(55)

    tree.delete(13)
    tree.delete(18)
    fmt.Printf("tree:%+v\n", tree)

    tree.String()
}

/**
实现二叉查找树，支持插入、删除、查找
注意：是否存在相同的值
*/

type treeNode struct {
    key   int // 中序遍历的节点序号
    val   int
    left  *treeNode
    right *treeNode
}

func (p *treeNode) find(data int) treeNode {
    for p != nil {
        if p.val > data {
            p = p.left
        } else if p.val < data {
            p = p.right
        } else {
            return *p
        }
    }

    return treeNode{}
}

func (p *treeNode) insert(data int) *treeNode {
    if p == nil {
        p = &treeNode{val: data}
        return p
    }

    node := p
    for !reflect.DeepEqual(node, treeNode{}) {
        if data > node.val {
            if node.right == nil {
                node.right = &treeNode{val: data}
                return p
            }
            node = node.right
        } else if data < node.val {
            if node.left == nil {
                node.left = &treeNode{val: data}
                return p
            }
            node = node.left
        } else { // data = t.Val

        }
    }
    return p
}

func (p *treeNode) delete(data int) {
    var pp = &treeNode{} // pp记录的是p的父节点
    for p != nil && p.val != data {
        pp = p
        if data > p.val {
            p = p.right
        } else {
            p = p.left
        }
    }
    if p == nil {
        return // 没有找到
    }

    // 要删除的节点有两个子节点
    if p.left != nil && p.right != nil { // 查找右子树中最小节点
        minP := p.right
        minPP := p             // minPP表示minP的父节点
        for minP.left != nil { // 找到该节点的右子树最小节点
            minPP = minP
            minP = minP.left
        }
        p.val = minP.val // 将minP的数据替换到p中
        p = minP         // 下面就变成了删除minP了
        pp = minPP
    }

    // 删除节点是叶子节点或者仅有一个子节点
    child := &treeNode{}
    if p.left != nil {
        child = p.left
    } else if p.right != nil {
        child = p.right
    } else {
        child = nil
    }

    if pp == nil {
        p = child
    } else if pp.left == p {
        pp.left = child
    } else {
        pp.right = child
    }
}

// 前序遍历：根-左-右
func preOrder(root *treeNode) {
    if root == nil {
        return
    }
    fmt.Printf("node:%v\n", root)
    preOrder(root.left)
    preOrder(root.right)
}

// 中序遍历：左-根-右
func (p *treeNode) InOrderTraverse() {
    if p == nil {
        return
    }
    inOrderTraverse(p)
}

func inOrderTraverse(root *treeNode) {
    if root == nil {
        return
    }
    inOrderTraverse(root.left)
    fmt.Printf("node.val:%v\n", root.val)
    inOrderTraverse(root.right)
}

func postOrder(root *treeNode) {
    if root == nil {
        return
    }
    inOrderTraverse(root.left)
    inOrderTraverse(root.right)
    fmt.Printf("node:%v\n", root)
}

func (p *treeNode) String() {
    if p == nil {
        fmt.Println("Tree is empty!")
        return
    }
    stringify(p, 0)
    println("---------------")
}

func stringify(node *treeNode, level int) {
    if node == nil {
        return
    }
    format := ""
    for i := 0; i < level; i++ {
        format += "\t" // 根据节点的深度决定缩进长度
    }
    format += "---[ "
    level++
    // 先递归打印左子树
    stringify(node.left, level)
    // 打印值
    fmt.Printf(format+"%d\n", node.val)
    // 再递归打印右子树
    stringify(node.right, level)
}

/**
面试题：https://leetcode-cn.com/problems/binode-lcci/
把二叉搜索树转换为单向链表
*/
func (p *treeNode) convertBiNode() *treeNode {
    return nil
}

// path Sum路径总和 go语言实现(DFS)
// 题解：https://leetcode-cn.com/problems/path-sum/solution/lu-jing-zong-he-de-si-chong-jie-fa-dfs-hui-su-bfs-/
// 时间
func hasPathSum(root *treeNode, sum int) bool {
    if root==nil {
        return false
    }
    if root.left == root.right {
        return sum == root.val
    }
    return hasPathSum(root.left,sum-root.val) || hasPathSum(root.right,sum-root.val)
}

func hasPathSumBFS(root *treeNode, sum int) bool {
    if root == nil{
        return false
    }
    queueNode:=[]*treeNode
    queueVal :=[]int{}
    queueNode = append(queueNode,root)
    queueVal = append(queueVal,root.val)
    for len(queueNode)!=0 {
        now:=queueNode[0]
        queueNode = queueNode[1:]
        temp:=queueVal[0]
        queueVal:=queueVal[1:]
        if now.left == nil && now.right == nil {
            if temp == sum {
                return true
            }
            continue
        }
        if now.left!=nil {
            queueNode = append(queueNode,now.left)
            queueVal = append(queueVal,now.left.val+temp)
        }
        if now.right!=nil {
            
        }
    }
}

/**
Invert Binary Tree (翻转二叉树)
递归采用：深度优先搜索，前序遍历
*/
func invertTree(root *treeNode) *treeNode {
    if root == nil {
        return root
    }

    // 当前需要一级是把左节点指向右节点，右节点指向已经翻转的左节点
    root.left, root.right = root.right, root.left
    invertTree(root.left)
    invertTree(root.right)
    return root
}

// 最大深度
func maxDepth(root *treeNode) int {
    if root == nil {
        return 0
    }

    if root.left == nil && root.right == nil {
        return 1
    }
    return int(math.Max(float64(maxDepth(root.left)), float64(maxDepth(root.right)))) + 1
}
