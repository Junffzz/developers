package syncs

// import (
//     "math"
//     "sync"
//     "sync/atomic"
//     "unsafe"
// )
//
// // 用于表示并发安全的字典的接口
// type ConcurrentMap interface {
//     // 用于返回并发量
//     Concurrency() int
//     // Put会推送一个键-元素对
//     // 注意！参数element的值不能为nil
//     // 第一个返回值表示十分新增了键值对
//     // 若键已存在，新元素值会替换旧的元素值
//     Put(key string, element interface{}) (bool, error)
//     // Get会获取与指定键关联的那个元素
//     // 若返回nil，则说明指定的键不存在
//     Get(key string) interface{}
//     // 若结果值为true，则说明键已存在且删除，否则说明键不存在
//     Delete(key string) bool
//     // Len会返回当前字典中键值对的数量
//     Len() uint64
// }
//
// // 用于表示ConcurrentMap接口的实现类型
// type myConcurrentMap struct {
//     concurrency int       //表示并发量，同时也代表了segments字段的长度
//     segments    []Segment // 一个Segment类型值代表一个散列段。每个散列段都提供对其包含的键值对的读写操作。
//     total       uint64
// }
//
// // 创建一个ConcurrentMap类型的实例
// // 参数pairRedistributor可以为Nil
// func NewConcurrentMap(concurrency int, pairRedistributor PairRedistributor) ConcurrentMap {
//     if concurrency <= 0 {
//         return nil
//     }
//     if concurrency > MAX_CONCURRENCY {
//         return nil
//     }
//     cmap := &myConcurrentMap{}
//     cmap.concurrency = concurrency
//     cmap.segments = make([]Segment, concurrency)
//     for i := 0; i < concurrency; i++ {
//         cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
//     }
//     return cmap
// }
//
// func (cmap *myConcurrentMap) Put(key string, element interface{}) (bool, error) {
//     p := newPair(key, element)
//     s := cmap.findSegment(p.Hash())
//     ok, err := s.Put(p)
//     if ok {
//         atomic.AddUint64(&cmap.total, 1)
//     }
//     return ok, err
// }
//
// // 用于根据给定参数寻找并返回对应散列段
// func (cmap *myConcurrentMap) findSegment(keyHash uint64) Segment {
//     if cmap.concurrency == 1 {
//         return cmap.segments[0]
//     }
//     var keyHash32 uint32
//     if keyHash > math.MaxUint32 {
//         keyHash32 = uint32(keyHash >> 32)
//     } else {
//         keyHash32 = uint32(keyHash)
//     }
//     return cmap.segments[int(keyHash32>>16)%(cmap.concurrency-1)]
// }
//
// func (cmap *myConcurrentMap) Get(key string) interface{} {
//     keyHash := hash(key)
//     s := cmap.findSegment(keyHash)
//     pair := s.GetWithHash(key, keyHash)
//     if pair == nil {
//         return nil
//     }
//     return pair.Element()
// }
//
// func (cmap *myConcurrentMap) Delete(key string) bool {
//     s := cmap.findSegment(hash(key))
//     if s.Delete(key) {
//         atomic.AddInt64(&cmap.total, ^uint64(0))
//         return true
//     }
//     return false
// }
//
// // 用于表示并发安全的键值对的接口
// type Pair interface {
//     // 单链键值对接口
//     linkedPair
//     // 返回键的值
//     Key() string
//     // 返回键的散列值
//     Hash() uint64
//     //  返回元素的值
//     Element() interface{}
//     // 设置元素的值
//     SetElement(element interface{}) error
//     // 生成一个当前键值对的副本并返回
//     Copy() Pair
//     // 返回当前键值对的字符串表示形式
//     String() string
// }
//
// // 用于表示单向链接的键值对的接口
// type linkedPair interface {
//     // Next用于获得下一个键值对
//     // 若返回值为nil，则说明当前已在单链表的末尾
//     Next() Pair
//     // SetNext 用于设置下一个键值对
//     // 这样就可以形成一个键值对的单链表
//     SetNext(nextPair Pair) error
// }
//
// // 用于表示键值对的类型
// type pair struct {
//     key string
//     // 用于表示键的散列值
//     hash    uint64
//     element unsafe.Pointer
//     next    unsafe.Pointer
// }
//
// // 用于创建一个Pair类型的实例
// func newPair(key string, element interface{}) Pair {
//     p := &pair{
//         key:  key,
//         hash: hash(key),
//     }
//     if element == nil {
//         return nil
//     }
//     p.element = unsafe.Pointer(&element)
//     return p
// }
//
// // 用于表示并发安全的散列段的接口
// type Segment interface {
//     // Put会根据参数放入一个键值对
//     // 第一个返回值表示是否新增了键值对
//     Put(p Pair) (bool, error)
//     // Get会根据给定参数返回对应的键值对
//     // 该方法会根据给定的键计算散列值
//     Get(key string) Pair
//     // GetWithHash会根据给定参数返回对应的键值对
//     // 注意！参数keyHash是基于参数key计算得出的散列值
//     GetWithHash(key string, keyHash uint64) Pair
//     // Delete会删除指定键的键值对
//     // 若返回值为true则说明已删除，否则说明未找到该键
//     Delete(key string) bool
//     // 获取当前段的尺寸（其中包含的散列桶的数量）
//     Size() uint64
// }
//
// // 用于表示并发安全的散列段的类型
// type segment struct {
//     // 用于表示散列桶切片
//     buckets []Bucket
//     // 用于表示散列桶切片的长度
//     bucketsLen int
//     // 用于表示键值对总数
//     pairTotal uint64
//     // 用于表示键值对的再分布器
//     pairRedistributor PairRedistributor
//     lock              sync.Mutex
// }
//
// func newSegment(bucketNumber int, pairRedistributor PairRedistributor) Segment {
//     if bucketNumber <= 0 {
//         bucketNumber = DEFAULT_BUCKET_NUMBER
//     }
//     if pairRedistributor == nil {
//         pairRedistributor = newDefaultPairRedistributor(DEFAULT_BUCKET_LOAD_FACTOR, bucketNumber)
//     }
//     buckets := make([]Bucket, bucketNumber)
//     for i := 0; i < bucketNumber; i++ {
//         buckets[i] = newBucket()
//     }
//     return &segment{
//         buckets:           buckets,
//         bucketsLen:        bucketNumber,
//         pairRedistributor: pairRedistributor,
//     }
// }
