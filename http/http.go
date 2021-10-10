package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

var (
    c = &http.Client{
        Timeout: 10*time.Second, //优先级最高
    }
)

func main() {
    // 控制超时时间，适合中断场景
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Nanosecond)
    defer cancel()

    req, err := http.NewRequest("GET", "https://mengkang.net/1405.html", nil)
    if err != nil {
        log.Fatal(err)
        return
    }
    req = req.WithContext(ctx)
    resp, err := c.Do(req)
    if err != nil {
        log.Fatal(err)
        return
    }
    defer resp.Body.Close()
    b, _ := ioutil.ReadAll(resp.Body)

    fmt.Printf("body:%s\n", b)
}
