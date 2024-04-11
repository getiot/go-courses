package main

import (
    "errors"
    "fmt"
)

func main() {
    err := doSomething1()
    if err != nil {
        goto HandleError
    }

	err = doSomething2()
    if err != nil {
        goto HandleError
    }

    fmt.Println("Success")
    return

HandleError:
    fmt.Println("Error occurred:", err)
    // 错误处理代码
}

func doSomething1() error {
    // 模拟一个错误
    return errors.New("something 1 went wrong")
}

func doSomething2() error {
    // 模拟一个错误
    return errors.New("something 2 went wrong")
}