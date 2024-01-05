package main

import "fmt"

func sol(arr []int) {
	var ans []int
	
	for _, v := range arr {
		if v >= 50 && v % 2 == 0 {
			ans = append(ans, v / 2)
		} else if v < 50 && v % 2 == 1 {
			ans = append(ans, v * 2)
		} else {
			ans = append(ans, v)
		}
	}
	fmt.Println(ans)
}

func main() {
    arr := []int{1, 2, 3, 100, 99 ,98}
    sol(arr)
}