package main

import "fmt"

func sol(num_list []int, n int) {
	fmt.Println(num_list[n-1:])
	
}

func main() {
	s1 := []int{2, 1, 6}
	n := 3
	sol(s1, n)
}