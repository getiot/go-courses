package main

import "fmt"

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("i 是一个整数")
	case float64:
		fmt.Println("i 是一个浮点数")
	case string:
		fmt.Println("i 是一个字符串")
	default:
		fmt.Println("i 的类型未知")
	}
}

func main() {
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}
