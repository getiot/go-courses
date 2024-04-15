package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

var index int = 1

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
    printBook(&book1)
    printBook(&book2)
    fmt.Printf("-------+-------------------------------------------------------------\n")
}

func printBook(book *Books) {
    fmt.Printf("Book %d | %10s | %10s | %12s | %10d\n", index, book.title, book.author, book.subject, book.book_id)
    index++
}