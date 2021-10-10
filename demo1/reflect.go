package demo1

import (
    "encoding/json"
    "fmt"
    "reflect"
)

var str = `{"uid":"uid","data":"{\"a\":1,\"b\":\"2\"}"}`

type Task struct {
    Uid   string      `json:"uid"`
    Data interface{}   `json:"data"`
}

func (t *Task) Insert()  {

}

type Args struct {
    A int       `json:"a"`
    B string    `json:"b"`
}

func newTask(args Args) *Task {
    t := new(Task)
    t.Data = args
    return t
}

func T(args Args) error {
    fmt.Println(args.A, args.B, "in")
    return nil
}

func main() {

    var t Task
    err := json.Unmarshal([]byte(str), &t)
    if err != nil {
        panic(err)
    }

    fType := reflect.TypeOf(T)
    if fType.Kind() != reflect.Func {
        panic("xx")
    }

    paramValue := reflect.New(fType.In(0))
    err = json.Unmarshal([]byte(t.Data.(string)), paramValue.Interface())
    fValue := reflect.ValueOf(T)
    fValue.Call([]reflect.Value{paramValue.Elem()})
}