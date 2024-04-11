package main

import "fmt"

func main() {
	var array = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i, x := range array {
		fmt.Printf("[%d] %f\n", i, x)
	}
}