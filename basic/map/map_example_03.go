package main

import "fmt"

func main() {
    var scores map[string]int

    /* 初始化 */
    scores = map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Carol": 95,
    }

    // scores["David"] = 80

    checkKey(scores, "Alice")
    checkKey(scores, "David")
}

func checkKey(scores map[string]int, name string) {

    /* 查看元素在集合中是否存在 */
    score, ok := scores[name]

    /* 如果 ok 是 true, 则存在，否则不存在 */
    if (ok) {
        fmt.Printf("%s's score is %d\n", name, score)
    } else {
        fmt.Printf("%s doesn't exist\n", name) 
    }
}

func isKeyExisted(scores map[string]int, name string) bool {
    _, ok := scores[name]
    return ok
}