package main

import (
	"strconv"
)

func sol(a,b int) int{
	ab, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	ba, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))
	
	if ab >= ba {
		return ab
	} else {
		return ba
	}
}
	
func main() {
    sol(98, 4)
}