package main

import (
    "fmt"
    "reflect"
)

func main() {
    var a float32 = 6
    var b bool = true
    var c *int
    d := "getiot.tech"
    e := make(map[string]string)

    fmt.Println(reflect.ValueOf(a))
    fmt.Println(reflect.ValueOf(b))
    fmt.Println(reflect.ValueOf(c))
    fmt.Println(reflect.ValueOf(d))
    fmt.Println(reflect.ValueOf(e))
    fmt.Println(reflect.ValueOf(nil))
}