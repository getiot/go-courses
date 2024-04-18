package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()
	fmt.Println("欢迎来到 GetIoT.tech！")
	time.Sleep(time.Second)
}
