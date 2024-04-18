package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x  = 0
	wg sync.WaitGroup
	// lock sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	// lock.Unlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	// lock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
	// lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}