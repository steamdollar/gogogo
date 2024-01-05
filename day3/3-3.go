package main

import "fmt"

func sol (my_string string, k int) string {
	var ans string
	for i :=0; i < k; i++ {
		ans += my_string
	}
	
	fmt.Println(ans)
	return ans
}

func main() {
	sol("string", 3)
}