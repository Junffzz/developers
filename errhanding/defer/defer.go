/*
@Time : 2020/10/28 09:37
@Author : ZhaoJunfeng
@File : defer
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func main()  {
	var err error
	tryDefer()
	err.Error()
}

func writeFile(filename string)  {
	file,err:=os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	if err!=nil {
		if pathError,ok:=err.(*os.PathError);!ok {
			panic(err)
		} else{
			fmt.Println(pathError.Op,pathError.Path,pathError.Err)
		}
		fmt.Println("Error:",err)
	}
	defer file.Close()

	write:=bufio.NewWriter(file)
	defer write.Flush()

	//f:=fib.Fibonacci()
	//for i:=0;i<20;i++ {
	//	fmt.Fprintln(write,f())
	//}
}
