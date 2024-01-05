package main

import (
	"fmt"
	"sort"
)

func sol(arr []int) []int {
	// 오름차순 정렬
	sort.Ints(arr)
	return arr[5:]
	
}

func main() {
    a := []int{12, 4, 15, 46, 38, 1, 14, 56, 32, 10}
    fmt.Println(sol(a))
}