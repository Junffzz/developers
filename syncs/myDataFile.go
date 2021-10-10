package syncs

import (
    "io"
    "os"
    "sync"
    "sync/atomic"
)

/**
备注：
1.锁实现在内部都用到了另一种同步机制———信号量
2.条件变量总要与互斥量组合使用。sync.NewCond
*sync.Cond类型的方法集合中有3个方法，即：Wait、Signal和Broadcast，分别代表了等待通知、单发通知和广播通知
*/

// 用于表示数据的类型
type Data []byte

// 用于表示数据文件的接口类型
type DataFile interface {
    // 读取一个数据块
    Read() (rsn int64, d Data, err error)
    // 写入一个数据块
    Write(d Data) (wsn int64, err error)
    // 获取最后读取的数据块的序列号，Reading Serial Number的缩写
    RSN() int64
    // 获取最后写入的数据块的序列号,Writing Serial Number的缩写
    WSN() int64
    // 获取数据块的长度
    DataLen() uint32
    // 关闭数据文件
    Close() error
}

// 用于表示数据文件的实现类型
type myDataFile struct {
    f       *os.File     // 文件
    rcond   *sync.Cond   // 条件变量
    fMutex  sync.RWMutex // 用于文件的读写锁
    wOffset int64        // 写操作需要用到的偏移量
    rOffset int64        // 读操作需要用到的偏移量
    //wMutex  sync.Mutex   // 写操作需要用到的互斥锁
    //rMutex  sync.Mutex   // 读操作需要用到的互斥锁
    dataLen uint32       // 数据块长度
}

func NewDataFile(path string, dataLen uint32) DataFile {
    f, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    if dataLen == 0 {
        panic("Invalid data length!")
    }

    df := &myDataFile{f: f, dataLen: dataLen}

    // 该rcond与fMutex的读锁关联的。
    df.rcond = sync.NewCond(df.fMutex.RLocker())
    return df
}

/**
 * 1.获取并更新读偏移量
2.根据读偏移量从文件中读取一块数据；
3.把该数据块封装成一个Data类型值并将其作为结果值返回
 * @Description:
 * @Params:
 * @date: 2021/2/13
*/
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
    // 读取并更新读偏移量
    var offset int64
    //df.rMutex.Lock()
    //offset = df.rOffset
    //df.rOffset += int64(df.dataLen)
    //df.rMutex.Unlock()

    // 改为原子CAS操作
    for {
        offset = atomic.LoadInt64(&df.rOffset)
        if atomic.CompareAndSwapInt64(&df.rOffset, offset, (offset + int64(df.dataLen))) {
            break
        }
    }

    // 读取一个数据块
    rsn = offset / int64(df.dataLen)
    bytes := make([]byte, df.dataLen)

    df.fMutex.RLock()
    defer df.fMutex.RUnlock()
    /*
       假如3个goroutine并发执行Read方法，2个goroutine并发执行Write方法。写的goroutine比读的goroutine少。
       过不了多久，读偏移量rOffset会等于甚至大于写偏移量wOffset的值。
    */
    for {
        _, err = df.f.ReadAt(bytes, offset)
        if err != nil {
            // io.EOF是error类型，所以这个err看成一种边界情况
            if err == io.EOF {
                // 暂时"放弃"fmutex的读锁并等待通知的到来，也就意味着Write方法中的数据块写操作不会受它的阻碍
                df.rcond.Wait() // 调用前需要锁定与之关联的读锁。此时，该goroutine处于阻塞状态
                continue
            }
            return
        }
        d = bytes
        return
    }
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
    // 读取并更新写偏移量
    var offset int64
    //df.wMutex.Lock()
    //offset = df.wOffset
    //df.wOffset += int64(df.dataLen)
    //df.wMutex.Unlock()

    // 改为原子CAS操作
    for {
        offset = atomic.LoadInt64(&df.wOffset)
        if atomic.CompareAndSwapInt64(&df.wOffset, offset, (offset + int64(df.dataLen))) {
            break
        }
    }

    // 写入一个数据块
    wsn = offset / int64(df.dataLen)
    var bytes []byte
    if len(d) > int(df.dataLen) {
        bytes = d[0:df.dataLen]
    } else {
        bytes = d
    }
    df.fMutex.Lock()
    defer df.fMutex.Unlock()
    _, err = df.f.Write(bytes)
    // 发送rcond的通知，让Read()方法可以执行
    df.rcond.Signal()
    return
}

func (df *myDataFile) RSN() int64 {
    //df.rMutex.Lock()
    //defer df.rMutex.Unlock()
    //return df.rOffset / int64(df.dataLen)

    // 原子操作写法
    offset := atomic.LoadInt64(&df.rOffset)
    return offset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
    //df.wMutex.Lock()
    //defer df.wMutex.Unlock()
    //return df.wOffset / int64(df.dataLen)

    // 原子操作写法
    offset := atomic.LoadInt64(&df.wOffset)
    return offset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
    return df.dataLen
}

func (df *myDataFile) Close() error {
    return nil
}
