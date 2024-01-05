package main

import (
	"fmt"
	"strconv"
)

func sol(arr []int) int {
	var even, odd string
	
	for _, v := range arr {
		if v % 2 != 0 {
			odd += strconv.Itoa(v)
		} else {
			even += strconv.Itoa(v)
		}
	}
	
	fmt.Println(odd, even)
	
	addOdd, _ := strconv.Atoi(odd)
	addEven, _ := strconv.Atoi(even)
	
	return addOdd + addEven
	
}

func main() {
	arr := []int{3,4,5,2,1}
    	fmt.Println(sol(arr))
}