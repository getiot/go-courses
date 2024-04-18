package main

import "fmt"

func receive(ch chan int) {
	for i := range ch {
		fmt.Printf("v:%v", i)
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	receive(ch)
}