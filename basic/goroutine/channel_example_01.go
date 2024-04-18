package main

import "fmt"

func main() {
    // ch := make(chan int)
	ch := make(chan int, 1)
    ch <- 100
    fmt.Println("发送成功")
}