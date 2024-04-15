package main

import "fmt"

func main() {
    ages := map[string]int{"Alice": 30, "Bob": 25, "Carol": 35}
    
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}