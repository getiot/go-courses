package main

import "fmt"

func receive(ch chan int) {
    ret := <-ch
    fmt.Println("接收成功", ret)
}

func main() {
    ch := make(chan int)
    go receive(ch)
    ch <- 100
    fmt.Println("发送成功")
}