package main

import "fmt"

func sol(n int) int {
	var ans int
	if n % 2 != 0 {
		for i := n; i > 0; i = i - 2 {
			ans += i
		}
		
		return ans
	} else {
		
		for i := n; i > 0; i = i - 2 {
			ans += i*i
		}
		return ans
	}
	
}

func main() {
	fmt.Println(sol(10))
}