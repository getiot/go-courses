package main

import "fmt"

func main() {
    var a int= 20   /* 声明实际变量 */
    var ptr *int     /* 声明指针变量 */

    ptr = &a  /* 指针变量的存储地址 */

    fmt.Printf("  a 变量的内存地址: %x\n", &a)

    /* 指针变量的存储地址 */
    fmt.Printf("ptr 变量的存储地址: %x\n", ptr)

    /* 使用指针访问值 */
    fmt.Printf("*ptr 变量的值: %d\n", *ptr)
}