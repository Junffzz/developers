/*
@Time : 2020/12/23 14:02
@Author : ZhaoJunfeng
@File : stack_test
*/
package main

import (
	"fmt"
	"testing"
)

/**
string 大、中、小括号=》合法？
1."()"
2."()[]"
3."([)]"
4."((([])))"
5."]][["
*/
func TestBracketValid(t *testing.T) {
	if ok := IsValid("()"); ok == true {
		fmt.Printf("TestBracketValid is ok!\n")
	} else{
		fmt.Printf("TestBracketValid is fail!\n")
	}
}
