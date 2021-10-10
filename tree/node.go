/*
@Time : 2020/10/20 09:12
@Author : ZhaoJunfeng
@File : node
*/
package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Print(node.value," ")
}

func (node *treeNode) traverse()  {
	if node==nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	root.setValue(100)

	pRoot:=&root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()
}

//扩展已有类型

type myTreeNode struct {
	node *treeNode
}

func (myNode *myTreeNode) postOrder()  {
	if myNode==nil {
		return
	}

	leftNode := &myTreeNode{myNode.node.left}
	leftNode.postOrder()

	rightNode := &myTreeNode{myNode.node.left}
	rightNode.postOrder()

	myNode.node.print()
}