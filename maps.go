/*
@Time : 2020/10/16 10:03
@Author : ZhaoJunfeng
@File : maps
*/
package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2 ==empty map
	var m3 map[string]int      //m3==nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting Value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	courseName, ok = m["cours"] //key不存在，会打印出空串
	fmt.Println(courseName, ok)

	fmt.Println("Deleting values")
	name,ok:=m["name"]
	fmt.Println(name,ok)

	delete(m,"name")
	name,ok=m["name"]
	fmt.Println(name,ok)


}
