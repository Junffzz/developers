/*
@Time : 2021/4/21 10:26
@Author : ZhaoJunfeng
@File : growslice
*/
package main

import "fmt"

type UserData struct {
    Name  string
}

func main() {
    // tUnsafe()
    // var a,i int
    // for i = 0; i < 10; i++ {
    //     a+=i
    // }

    var info UserData
    info.Name = "WilburXu"
    _ = GetUserInfo(&info)
}

func tUnsafe()  {
    var a,i int
    for i = 0; i < 10; i++ {
        a+=i
    }
    fmt.Printf("i:%v,a:%v\n",i,a)
    return
}

func GetUserInfo(userInfo *UserData) *UserData {
    return userInfo
}