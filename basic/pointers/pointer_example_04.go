package main

import "fmt"

const MAX int = 3

func main() {
    a := []int{10,100,200}

    for i := 0; i < MAX; i++ {
        fmt.Printf("a[%d] = %d\n", i, a[i])
    }
}