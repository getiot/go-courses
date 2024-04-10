package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var score int = 90

	switch score {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B":
		fmt.Printf("良好\n")
	case grade == "C":
		fmt.Printf("及格\n")
	default:
		fmt.Printf("不及格\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
