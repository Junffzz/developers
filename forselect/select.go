package main

import (
    "fmt"
    "time"
)

func main() {
    var t chan int
    t = make(chan int,1)
    t <-1
    go func() {
        var end = false
        for {
            select {
            case _ = <-t:
                fmt.Printf("0.select t.\n")
                end=true
            }
            if end==true {
                fmt.Printf("1.for t.\n")
                break
            }

        }
        fmt.Printf("2.for t.\n")
    }()

    time.Sleep(3 * time.Second)
}
