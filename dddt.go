/*
@Time : 2021/4/12 20:57
@Author : ZhaoJunfeng
@File : dddt
*/
package main

import "fmt"

func main() {
    s := "tabcballabcdcba"
    var Slist = make([]string, 0, len(s))

    n := len(s)
    for i := range s {
        for j := i; j < n; j++ {
            Slist = append(Slist, s[i:j])

        }
    }

    var res string
    max := 0
    for h := range Slist {
        if ok := isHWs(Slist[h]); ok {
            if yes:=isXmax(len(Slist[h]),max);yes {
                max = len(Slist[h])
                res = Slist[h]
            }
        }
    }
    fmt.Printf("res :%s",res)
}

func isHWs(s string) bool {
    n := len(s)
    if n == 0 {
        return false
    }
    for i := 0; i <= n/2; i++ {
        if s[i] != s[n-i+1] {
            return false
        }
    }
    return true
}

func isXmax(x, y int) bool {
    if x > y {
        return true
    }
    return false
}
