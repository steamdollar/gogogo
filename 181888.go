package main

import "fmt"

func sol(arr []int, n int) []int {
	var ans []int
	for i := 0; i < len(arr); i = i + n {
		ans = append(ans, arr[i])
	}
	
	return ans
}

func main() {
	
	arr := []int{4, 2, 6, 1, 7, 6}
    	fmt.Println(sol(arr, 2))
}