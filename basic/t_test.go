/*
@Time : 2021/10/8 15:10
@Author : ZhaoJunfeng
@File : t.go
*/
package basic

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "sync"
    "testing"
    "time"
)

func TestT1(t *testing.T) {
    var waitg sync.WaitGroup
    var urls = []string{"www.baidu.com", "www.sina.com", "www.xueersi.com", "www.qq.com", "www.aliyun.com"}
    httpClient := http.Client{
        Timeout: 3 * time.Second,
    }

    ctx := context.Background()
    ctx,cancel:=context.WithCancel(ctx)
    waitg.Add(len(urls))

    for _, url := range urls {
        go func(ctx context.Context, target string) {
            defer waitg.Done()

            req, err := http.NewRequestWithContext(ctx, "GET", "https://"+target, nil)
            if err != nil {
                t.Logf("http new request err:%v", err)
                return
            }
            resp, err := httpClient.Do(req)
            if err != nil {
                cancel()
                t.Logf("http err:%v\n", err)
                return
            }
            if target=="www.baidu.com" {
                cancel()
                return
            }
            ioutil.ReadAll(resp.Body)
            defer resp.Body.Close()
            // t.Logf("%s\n", body)

        }(ctx, url)
    }

    waitg.Wait()
    fmt.Println("finish!")
}
