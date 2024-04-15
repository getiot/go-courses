package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main() {
    var book1 Books  /* 声明 book1 为 Books 类型 */
    var book2 Books  /* 声明 book2 为 Books 类型 */

    /* book 1 描述 */
    book1.title = "Go 语言"
    book1.author = "getiot.tech"
    book1.subject = "Go 基础教程"
    book1.book_id = 978495407

    /* book 2 描述 */
    book2.title = "Python 语言"
    book2.author = "getiot.tech"
    book2.subject = "Python 基础教程"
    book2.book_id = 978495700

    /* 打印图书信息 */
    fmt.Printf("       |     title    |    author   |      subject     |   book_id\n")
    fmt.Printf("-------+-------------------------------------------------------------\n")
    fmt.Printf("Book 1 | %10s | %10s | %12s | %10d\n", book1.title, book1.author, book1.subject, book1.book_id)
    fmt.Printf("Book 2 | %10s | %10s | %12s | %10d\n", book2.title, book2.author, book2.subject, book2.book_id)
    fmt.Printf("-------+-------------------------------------------------------------\n")
}