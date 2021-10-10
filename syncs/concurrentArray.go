package syncs

import "sync/atomic"

// 并发安全的整数数组接口，该整型数组的长度是固定并且必须初始化时给定。
type ConcurrentArray interface {
    // 用于设置指定索引上的元素值
    Set(index uint32, elem int) (err error)
    // 用于获取指定索引上的元素值
    Get(index uint32) (elem int, err error)
    // 用于获取数组的长度
    Len() uint32
}

// 用于表示ConcurrentArray接口的实现类型
type concurrentArray struct {
    length uint32
    val    atomic.Value // 对ConcurrentArray接口的实现，必须是指针方法。因为原子值不该复制
}

// 创建一个ConcurrentArray 类型值
func NewConcurrentArray(length uint32) ConcurrentArray {
    array := concurrentArray{}
    array.length = length
    array.val.Store(make([]int, array.length))
    return &array
}

func (array *concurrentArray) Set(index uint32, elem int) (err error) {
    if err = array.checkIndex(index); err != nil {
        return
    }
    if err = array.checkValue(); err != nil {
        return
    }
    newArray := make([]int, array.length)
    copy(newArray, array.val.Load().([]int))
    newArray[index] = elem
    array.val.Store(newArray)
    return
}

func (array *concurrentArray) Get(index uint32) (elem int, err error) {
    if err = array.checkIndex(index); err != nil {
        return
    }
    if err = array.checkValue(); err != nil {
        return
    }
    elem = array.val.Load().([]int)[index]
    return
}

func (array *concurrentArray) checkIndex(index uint32) error {
    return nil
}

func (array *concurrentArray) checkValue() error {
    return nil
}

func (array *concurrentArray) Len() uint32 {
    return array.length
}