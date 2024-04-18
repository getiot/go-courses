package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("Hello")
	defer wg.Done() // 把计算器-1
}

func main() {
	wg.Add(1) // 把计数器+1
	go hello()
	fmt.Println("欢迎来到 GetIoT.tech！")
	wg.Wait() // 阻塞代码的运行，直到计算器为0
}
