package main

import "fmt"

func sol(s,e int) []int{
	var ans []int
	for i :=s; i <= e; i++ {
		ans = append(ans, i)
	}
	return ans
}

func main() {
    fmt.Println(sol(3, 10))
}