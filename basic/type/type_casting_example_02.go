package main
import "fmt"

func main() {  
    var a int64 = 3
    var b int32
    b = int32(a)
    fmt.Printf("b = %d\n", b)
}