/*
@Time : 2020/10/22 08:45
@Author : ZhaoJunfeng
@File : download
*/
package main

import (
	"devtest/retiever/mock"
	"devtest/retiever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func post(poster Poster)  {
	poster.Post("www.imooc.com", map[string]string{
		"name":"ccmouse",
		"course":"golang",
	})
}

const url = "www.imooc.com"
//接口组合
type RetrieverPoster interface {
	Retriever
	Poster
	Connect(host string)
}

func assion(s RetrieverPoster)  {
	s.Get(url)
	s.Post(url,map[string]string{
		"name":"ccmouse",
		"course":"golang",
	})
}

func main() {
	//fmt.Sscanf()
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "test-zz",
		Timeout: time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
	realRetriever:=r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)
}

func inspect(r Retriever)  {
	fmt.Printf("%T %v\n",r,r)
	fmt.Println("Type switch")
	switch v:=r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case real.Retriever:
		fmt.Println("real userAgent:",v.UserAgent)

	}
}