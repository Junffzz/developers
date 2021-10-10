/*
@Time : 2020/7/10 09:40
@Author : ZhaoJunfeng
@File : debug1
*/
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRand(b *testing.B) {
	randS := rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		rand_t := randS.Intn(4) + 1
		fmt.Println("rand_t:", rand_t)
	}
}

