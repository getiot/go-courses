package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup // 等待组
)

// add 对全局变量x执行50000次加1操作
func add() {
	for i := 0; i < 50000; i++ {
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}