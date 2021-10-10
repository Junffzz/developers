/*
@Time : 2021/6/7 15:39
@Author : ZhaoJunfeng
@File : pointer
*/
package unsafes

import (
    "fmt"
    "testing"
    "unsafe"
)

// 类型转换
func TestPointer(t *testing.T) {
    num := 5
    numPointer := &num

    flnum := (*float32)(unsafe.Pointer(numPointer))
    fmt.Println(flnum)
    s:="123456"
    a:=int(s[2]-'0')
    fmt.Printf("result a: %v\n",a)
}

type Num struct {
    i string
    j int64
}

func TestOffsetof(t *testing.T) {
    n := Num{
        i: "EDDYCJY",
        j: 1,
    }
    nPointer := unsafe.Pointer(&n)

    niPointer := (*string)(unsafe.Pointer(nPointer))
    *niPointer = "煎鱼"

    njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
    *njPointer = 2

    fmt.Printf("n.i:%s, n.j:%d\n", n.i, n.j)

    ret := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
    fmt.Printf("bstFromPreorder:%v\n", ret)
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

func bstFromPreorder(preorder []int) *TreeNode {
    n := len(preorder)
    if n == 0 {
        return nil
    }
    root := &TreeNode{preorder[0], nil, nil}
    var i int
    for i = 1; i < n; i++ {
        if preorder[i] > root.Val {
            break
        }
    }
    root.Left = bstFromPreorder(preorder[1:i])
    if i < n {
        root.Right = bstFromPreorder(preorder[i:])
    }

    return root
}

