/*
@Time : 2020/10/21 11:28
@Author : ZhaoJunfeng
@File : errors
*/
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	//var err error
	//err = errors.Wrapf(err, "zjftest %v", "eeee")
	//fmt.Println(err)
	//fmt.Println(fmt.Errorf("---debug:%w", err))

	err := errors.New("whoops")

	err = test1()
	if err != nil {
		err = errors.Wrap(err, "zjftest.")
		err = errors.Wrap(err, "oh noes")
		//err = fmt.Errorf("zjftest.%w",err)
		//err = fmt.Errorf("oh noes.%w",err)
	}

	//fmt.Printf("%v\n", err)
	fmt.Printf("%+v", err)
}

func test1() error {

	return fmt.Errorf("insert data is nil")
}
