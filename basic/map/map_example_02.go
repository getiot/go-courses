package main

import "fmt"

func main() {
    var scores map[string]int

    /* 创建集合 */
    scores = make(map[string]int)

    /* 添加初始数据 */
    scores = map[string]int{"Alice": 90, "Bob": 85, "Carol": 95}

    /* 遍历集合，使用 key 输出 map 值 */
    for name := range scores {
        fmt.Printf("%8s : %d\n", name, scores[name])
    }

    fmt.Println("----------------")

    /* 删除数据 */
    delete(scores, "Bob")
   
    /* 再次输出 */
    for name := range scores {
        fmt.Printf("%8s : %d\n", name, scores[name])
    }
}