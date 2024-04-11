package main

import "fmt"

func main() {
    /* 数组 - 3 行 4 列 */
    var array2d = [3][4]int {{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
    var i, j int

    /* 输出数组元素 */
    for  i = 0; i < 3; i++ {
        for j = 0; j < 4; j++ {
            fmt.Printf("a[%d][%d] = %d\n", i, j, array2d[i][j])
        }
    }
}