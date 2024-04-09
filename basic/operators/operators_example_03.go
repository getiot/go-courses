package main

import "fmt"

func main() {

	var a bool = true
	var b bool = false

	if a && b {
		fmt.Printf("a && b 为 true\n")
	} else {
		fmt.Printf("a && b 为 false\n")
	}

	if a || b {
		fmt.Printf("a || b 为 true\n")
	}

	if !(a && b) {
		fmt.Printf("!(a && b) 为 true\n")
	}
}
