package main

import "fmt"

func factorial(n int) (result int) {
    if n == 0 {
        result = 1   
    } else {
        result = n * factorial(n - 1)
    }
    return
}

func main() {  
    var i int = 5
    fmt.Printf("%d 的阶乘是 %d\n", i, factorial(i))
}