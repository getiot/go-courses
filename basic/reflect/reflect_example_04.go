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

    // 获取变量 person 的可设置的反射值，反射中使用 Elem() 方法获取指针所指向的值
    v := reflect.ValueOf(&person).Elem()

    v.FieldByName("Name").SetString("Tina") // 设置 Name 字段的值
    v.FieldByName("Age").SetInt(18)         // 设置 Age 字段的值

    fmt.Println(v.FieldByName("Name").String()) // 输出名字
    fmt.Println(v.FieldByName("Age").Int())     // 输出年龄
}