package main

import "fmt"

func main() {
	ch := make(chan int, 1) // 创建一个类型为int，缓冲区大小为1的通道

	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch: // 第一次循环由于没有值，所以该分支不满足
			fmt.Println(x)
		case ch <- i: // 将i发送给通道(由于缓冲区大小为1，缓冲区已满，第二次不会走该分支)
		}
	}
}
