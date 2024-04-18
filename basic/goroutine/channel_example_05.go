package main

import "fmt"

func main() {
    ch := make(<-chan int, 1)  // 这个是错的
    ch <- 100
    fmt.Println("发送成功")
}