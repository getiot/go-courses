package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("%d + %d 等于 %d\n", a, b, c)
	c = a - b
	fmt.Printf("%d - %d 等于 %d\n", a, b, c)
	c = a * b
	fmt.Printf("%d * %d 等于 %d\n", a, b, c)
	c = a / b
	fmt.Printf("%d / %d 等于 %d\n", a, b, c)
	c = a % b
	fmt.Printf("%d %% %d 等于 %d\n", a, b, c)
	a++
	fmt.Printf("a++ 等于 %d\n", a)
	a--
	fmt.Printf("a-- 等于 %d\n", a)
}
