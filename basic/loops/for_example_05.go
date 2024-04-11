package main

import "fmt"

func main() {
    /* 定义局部变量 */
    var a int = 10

    /* for 循环 */
    for a < 20 {
        if a == 15 {
            /* 跳过此次循环 */
            a = a + 1
            continue
        }
        fmt.Printf("%d\n", a)
        a++;     
     }  
}