package main
/**
题目介绍：https://www.nowcoder.com/practice/e3769a5f49894d49b871c09cadd13a61?tpId=117&tqId=1006010&tab=answerKey
 * 题目描述
设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值
[要求]
set和get方法的时间复杂度为O(1)
某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
对于每个操作2，输出一个答案
 * @Description: 
 * @Params: 
 * @date: 2021/2/9 
 */

func main() {

}


/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU( operators [][]int ,  k int ) []int {
    // write code here
    return nil
}

/**
https://leetcode-cn.com/problems/lru-cache-lcci/solution/goshuang-xiang-lian-biao-map-shi-xian-lru-by-pengt/
 */
type LinkNode struct{
    key, value int
    pre, next *LinkNode
}

type LRUCache struct {
    m map[int]*LinkNode
    capacity int
    head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
    head := &LinkNode{-1, -1, nil, nil}
    tail := &LinkNode{-1, -1, nil, nil}
    head.next = tail
    tail.pre = head
    cache := LRUCache{make(map[int]*LinkNode), capacity, head, tail}
    return cache
}

func (this *LRUCache) AddNode(node *LinkNode) {
    node.pre = this.head
    node.next = this.head.next
    this.head.next = node
    node.next.pre = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
    node.pre.next = node.next
    node.next.pre = node.pre
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
    this.RemoveNode(node)
    this.AddNode(node)
}

func (this *LRUCache) Get(key int) int {
    m := this.m
    if node, ok := m[key]; ok {
        this.MoveToHead(node)
        return node.value
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int)  {
    m := this.m
    if node, ok := m[key]; ok {
        node.value = value
        this.MoveToHead(node)
    } else {
        n := &LinkNode{key, value, nil, nil}
        if len(m) >= this.capacity {
            delete(m, this.tail.pre.key)
            this.RemoveNode(this.tail.pre)
        }
        m[key] = n
        this.AddNode(n)
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//作者：pengtuo
//链接：https://leetcode-cn.com/problems/lru-cache-lcci/solution/goshuang-xiang-lian-biao-map-shi-xian-lru-by-pengt/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。