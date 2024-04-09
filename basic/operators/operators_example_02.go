package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10

	fmt.Printf("a = %d, b = %d\n", a, b)

	if a == b {
		fmt.Printf("a 等于 b\n")
	} else {
		fmt.Printf("a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("a 小于 b\n")
	} else {
		fmt.Printf("a 不小于 b\n")
	}
	if a > b {
		fmt.Printf("a 大于 b\n")
	} else {
		fmt.Printf("a 不大于 b\n")
	}

	/* Lets change value of a and b */
	a = 5
	b = 20

	println("----------------")
	fmt.Printf("a = %d, b = %d\n", a, b)

	if a <= b {
		fmt.Printf("a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("b 大于等于 a\n")
	}
}
