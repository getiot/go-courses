package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50}
    
    for index, value := range numbers {
        fmt.Printf("index: %d, value: %d\n", index, value)
    }
}