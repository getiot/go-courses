package main

import "fmt"

func main() {

	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a 变量类型为 %T\n", a)
	fmt.Printf("b 变量类型为 %T\n", b)
	fmt.Printf("c 变量类型为 %T\n", c)
	fmt.Printf("ptr 变量类型为 %T\n", ptr)

	ptr = &a
	fmt.Printf("a = %d\n", a)
	fmt.Printf("ptr = %d\n", ptr)
	fmt.Printf("*ptr = %d\n", *ptr)
}
