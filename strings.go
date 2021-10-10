/*
@Time : 2020/10/18 22:19
@Author : ZhaoJunfeng
@File : strings
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网"
	fmt.Printf("%s\n", []byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	for i,ch:=range []rune(s) {
		fmt.Printf("(%d %c) ",i,ch)
	}

}
