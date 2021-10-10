/*
@Time : 2020/10/13 07:39
@Author : ZhaoJunfeng
@File : basic
*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main()  {
	euler()
}

//常量
func enums()  {
	const(
		cpp=iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp,python,golang,javascript)

	//b,kb,mb,gb,tb,pb
	const (
		b=1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func euler()  {
	c:=3+4i
	fmt.Println(cmplx.Abs(c))

	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi+1),
		//cmplx.Pow(math.E,1i*math.Pi)+1
		)

	//枚举
	enums()
}