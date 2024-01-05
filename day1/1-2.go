package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
//     fmt.Println("a =", a)
//     fmt.Println("b =", b)

    fmt.Printf("a = %v", a)
    fmt.Println(" ")
    fmt.Printf("b = %v", b)
}