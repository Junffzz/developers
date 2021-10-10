/*
@Time : 2020/11/6 17:26
@Author : ZhaoJunfeng
@File : maps
*/
package basic

import "fmt"

func main() {
	m := make(map[int]string, 3)
	fmt.Printf("m:%+v,len:%v\n", m,len(m))
}
