package main

import "fmt"

func main() {

    numbers := [6]int{4, 6, 2, 17}

    for i, x := range numbers {
        fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
    }
}