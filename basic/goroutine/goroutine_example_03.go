package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Printf("[%v] Hello, welcome to GetIoT.tech!\n", i)
	defer wg.Done() // goroutine 结束计数器-1
}

func main() {
	for i := 0; i < 5; i++ {
		go hello(i)
		wg.Add(1) // 启动一个goroutine计数器+1
	}
	wg.Wait() // 等待所有的goroutine执行结束
}
