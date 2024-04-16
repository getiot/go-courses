package main

import "fmt"

func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
    var i int

    for i = 0; i < 15; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
	fmt.Println()
}