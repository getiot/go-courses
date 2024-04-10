package main

import "fmt"

func main() {
	/* 局部变量定义 */
	var a int = 50

	/* 判断布尔表达式 */
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else if a < 40 {
		fmt.Printf("a 大于等于 20，小于 40\n")
	} else if a < 60 {
		fmt.Printf("a 大于等于 40，小于 60\n")
	} else {
		fmt.Printf("a 不小于 60\n")
	}
	fmt.Printf("a 的值为 %d\n", a)
}
