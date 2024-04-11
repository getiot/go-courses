package main

import "fmt"

func main() {
    var a, b = 10, 12

    /* for 循环 */
    for i := 0; i < 5; i++ {
        fmt.Printf("i 的值为 %d\n", i)
    }

    for a < b {
        fmt.Printf("a 的值为 %d\n", a)
        a++
    }
}