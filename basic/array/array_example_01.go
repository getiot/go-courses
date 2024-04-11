package main

import "fmt"

func main() {
	var array [10] int

	for i, x := range array {
		fmt.Printf("[%d] %d\n", i, x)
	}
}