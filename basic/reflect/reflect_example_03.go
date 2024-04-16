package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age int
}

func main() {
    person := Person{"Rudy", 30}

    v := reflect.ValueOf(person) // 获取变量 person 的反射值
    fmt.Println(v.FieldByName("Name").String()) // 输出名字
    fmt.Println(v.FieldByName("Age").Int())     // 输出年龄
}