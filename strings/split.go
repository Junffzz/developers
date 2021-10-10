/*
@Time : 2020/11/19 16:55
@Author : ZhaoJunfeng
@File : split
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s:=strings.Split("lecturepie.xueersi.com",".")
	fmt.Println(s[0],len(s))
}
