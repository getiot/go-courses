package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup
	m  sync.Mutex // 互斥锁
)

func add() {
	for i := 0; i < 50000; i++ {
		m.Lock()    // 修改前加锁
		x = x + 1
		m.Unlock()  // 修改后解锁
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