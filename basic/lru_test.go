/*
@Time : 2021/3/19 11:13
@Author : ZhaoJunfeng
@File : lru_test
*/
package basic

import (
    "fmt"
    "testing"
)

func TestLruCache(t *testing.T) {
    obj := Constructor(2)
    v1 := obj.Get(1)
    obj.Put(1, 13)
    v2 := obj.Get(1)

    fmt.Printf("v1:%v v2:%v\n", v1, v2)
}

type LRUCache struct {
    size       int
    capacity   int
    data       map[int]*dLinkedNode // 维护有序
    head, tail *dLinkedNode
}

type dLinkedNode struct {
    Key, Value int
    Pre, Next  *dLinkedNode
}

func Constructor(capacity int) LRUCache {
    p := LRUCache{
        capacity: capacity,
        data:     make(map[int]*dLinkedNode),
        head:     &dLinkedNode{},
        tail:     &dLinkedNode{},
    }
    p.head.Next = p.tail
    p.tail.Pre = p.head
    return p
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.data[key]; ok {
        this.moveToHead(node)
        return node.Value
    }
    return -1
}

/**
思考点：
1.先判断是否存在，如果存在要修改值，再返回
2.如果空间<capacity直接插入链表头部
3.如果链表空间满了，在链表尾部删除一个，然后在表头添加
*/
func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.data[key]; ok {
        node.Value = value
        this.moveToHead(node)
        return
    }

    node := &dLinkedNode{
        Key:   key,
        Value: value,
    }
    this.data[key] = node
    this.addToHead(node)
    this.size++
    // 满了
    if this.size > this.capacity {
        // tail淘汰
        removeNode := this.tail.Pre
        this.removeNode(removeNode)
        delete(this.data, removeNode.Key)
        this.size--
        return
    }
    return
}

func (this *LRUCache) addToHead(node *dLinkedNode) {
    node.Pre = this.head
    node.Next = this.head.Next
    this.head.Next.Pre = node
    this.head.Next = node
}

func (this *LRUCache) removeNode(node *dLinkedNode) {
    node.Pre.Next = node.Next
    node.Next.Pre = node.Pre
}

func (this *LRUCache) moveToHead(node *dLinkedNode) {
    // 如果该节点不是头节点，移动至头部
    this.removeNode(node)
    this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
