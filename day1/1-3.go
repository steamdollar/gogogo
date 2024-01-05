package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, a string
	fmt.Scan(&s1, &a)

	num, _ := strconv.Atoi(a)
	for i :=0; i < num; i++ {
		fmt.Print(s1)
	}
}