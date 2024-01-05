package main

import "fmt"

func sol(num, n, m int) int {
	if num % n == 0 && num % m == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(sol(10, 2, 5))
}