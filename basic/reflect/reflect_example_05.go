package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age int
}

func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
    person := Person{"Rudy", 30}

    v := reflect.ValueOf(person) // 获取变量 x 的反射值
    m := v.MethodByName("SayHello") // 获取名称为 MethodName 的方法

    if m.IsValid() { // 判断方法是否有效
        m.Call(nil) // 调用 SayHello 方法
    } else {
        fmt.Println("Method not found.")
    }
}