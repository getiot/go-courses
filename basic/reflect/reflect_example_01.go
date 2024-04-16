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

    fmt.Println(reflect.TypeOf(a))
    fmt.Println(reflect.TypeOf(b))
    fmt.Println(reflect.TypeOf(c))
    fmt.Println(reflect.TypeOf(d))
    fmt.Println(reflect.TypeOf(e))
    fmt.Println(reflect.TypeOf(nil))
}