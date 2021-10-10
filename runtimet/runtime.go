/*
@Time : 2020/12/17 19:10
@Author : ZhaoJunfeng
@File : runtime
*/
package runtimet

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func TestDir()  {
	GetDir()
}

func GetDir() {
	_, filename, _, _ := runtime.Caller(0)
	f, err := os.Open(path.Join(path.Dir(filename), "data.csv"))
	fmt.Printf("--------debug dir path:%v,f:%v,err:%w\n",path.Dir(filename),f,err)
}