package main

import "fmt"

func main() {

	var a int = 1
	var b int = 2
	var c int = 6
	var d int = 2
	var e int

	e = (a + b) * c / d
	fmt.Printf("   (a + b) * c / d 的值为 %d\n", e)

	e = ((a + b) * c) / d
	fmt.Printf(" ((a + b) * c) / d 的值为 %d\n", e)

	e = (a + b) * (c / d)
	fmt.Printf(" (a + b) * (c / d) 的值为 %d\n", e)

	e = a + (b * c) / d
	fmt.Printf("   a + (b * c) / d 的值为 %d\n", e)
}
