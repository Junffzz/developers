/*
@Time : 2021/5/26 10:07
@Author : ZhaoJunfeng
@File : lock_test
*/
package syncs

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

func TestLocker(t *testing.T) {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println(getItem())
        }()
    }
    wg.Wait()
}

type item struct {
    content string
}

var i *item
var once sync.Once

func getItem() *item {
    once.Do(func() {
        time.Sleep(time.Second * 10)
        i = &item{
            content: "set content",
        }
    })
    return i
}