package main

import "fmt"

func sol(arr []int) int {
	var sum int
	mul := 1
	for _, v := range arr {
		sum += v
		mul *= v
	}
	
	
	if mul > sum * sum {
		return 0
	} else {
		return 1
	}
}

func main() {
    fmt.Println(7)
    
    arr := []int{3, 4, 5, 2, 1}
    fmt.Println(sol(arr))
}