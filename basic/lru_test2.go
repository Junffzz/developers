/*
@Time : 2021/3/23 18:16
@Author : ZhaoJunfeng
@File : lru_test2
*/
package basic

// type LRUCache struct {
//     size       int
//     capacity   int
//     cache      map[int]*dLinkedNode
//     head, tail *dLinkedNode
// }
//
// type dLinkedNode struct {
//     key       int
//     value     int
//     pre, next *dLinkedNode
// }
//
// func Constructor(capacity int) LRUCache {
//     p := LRUCache{
//         size:     0,
//         capacity: capacity,
//         cache:    make(map[int]*dLinkedNode),
//         head:     &dLinkedNode{},
//         tail:     &dLinkedNode{},
//     }
//     p.head.next = p.tail
//     p.tail.pre = p.head
//     return p
// }
//
// func (this *LRUCache) Get(key int) int {
//     if node, ok := this.cache[key]; ok {
//         this.removeToHead(node)
//         return node.value
//     } else {
//         return -1
//     }
// }
//
// func (this *LRUCache) Put(key int, value int) {
//     if node, ok := this.cache[key]; ok {
//         node.value = value
//         this.removeToHead(node)
//         return
//     }
//     node := &dLinkedNode{
//         key:   key,
//         value: value,
//     }
//
//     this.size++
//     this.addToHead(node)
//     this.cache[key] = node
//     if this.size > this.capacity {
//         removeNode := this.tail.pre
//         this.removeNode(removeNode)
//         delete(this.cache, removeNode.key)
//     }
// }
//
// func (this *LRUCache) removeNode(node *dLinkedNode) {
//     node.pre.next = node.next
//     node.next.pre = node.pre
// }
//
// func (this *LRUCache) addToHead(node *dLinkedNode) {
//     node.pre = this.head
//     node.next = this.head.next
//     this.head.next.pre = node
//     this.head.next = node
// }
//
// func (this *LRUCache) removeToHead(node *dLinkedNode) {
//     this.removeNode(node)
//     this.addToHead(node)
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
