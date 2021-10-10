package main

import (
    "fmt"
    "reflect"
)

type Task struct {
    Uid  string      `json:"uid"`
    Data interface{} `json:"data"`
}

func (t *Task) Insert(data interface{}) {
    fmt.Printf("Insert is ok!\n")
}

func main() {
    var a = make(map[string]int)

    a1 := reflect.TypeOf(&a)
    a2 := a1.Elem()
    a3 := a2.Elem()
    fmt.Printf("a1:%v,a2:%v,a3:%v\n", a1, a2, a3)
    //xx := 1
    //c := &xx
    b := reflect.New(a3)
    b = reflect.ValueOf(2)

    //b2:=reflect.TypeOf(&b)
    fmt.Printf("b:%v,p:%v\n", b)
    c := &b
    cValue := reflect.ValueOf(c)
    f := reflect.ValueOf(test1)
    f.Call([]reflect.Value{cValue})

    // 测试调用结构体方法
    t := reflect.ValueOf(&Task{})

    s := reflect.TypeOf(&Task{}).Elem()
    s1 := reflect.New(s)
    s2 := s1.MethodByName("Insert")
    s2.Call([]reflect.Value{reflect.ValueOf("zjf")})

    setNameMethod := t.MethodByName("Insert")
    setNameMethod.Call([]reflect.Value{reflect.ValueOf("zjf")})
}

func test1(i interface{}) {
    i1 := reflect.TypeOf(i).Elem()
    fmt.Printf("i1:%v,i:%v\n", i1, i)
}
