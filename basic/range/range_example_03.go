package main

import "fmt"

func main() {
    str := "GetIoT"
    
    for index, char := range str {
        fmt.Printf("index: %d, char: %c\n", index, char)
    }
}