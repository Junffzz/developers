/*
@Time : 2020/5/20 19:29
@Author : ZhaoJunfeng
@File : test
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	jsoniter_extra "github.com/json-iterator/go/extra"
)

type Te struct {
	Tids []int `json:"cids"`
	Id int `json:"id"`
}

func init() {
	// json兼容性扩展
	jsoniter_extra.RegisterFuzzyDecoders()
}


func main()  {
	var test Te
	test = Te{
		Tids:[]int{12345678,321},
		Id:12345678,
	}
	a,_:=json.Marshal(test)
	d := json.NewDecoder(bytes.NewReader(a))
	d.UseNumber()
	var b map[string]interface{}
	d.Decode(&b)
	fmt.Printf("Decode b:%+v \n",b)
	fmt.Printf("Decode b.cids:%+v \n",b["cids"])
	m:=b["cids"]
	fmt.Printf("Decode b.cids :%+v \n",m)
	n := b["id"].(json.Number)
	fmt.Println(n.Float64())
	fmt.Println(n.Int64())
	fmt.Println(n.String())

	//json.Unmarshal
	json.Unmarshal(a,&b)
	fmt.Printf("Unmarshal b:%+v\n",b)

	jsoniter.Unmarshal(a,&b)
	fmt.Printf("jsoniter.Unmarshal b:%+v\n",b)
}