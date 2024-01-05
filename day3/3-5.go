package main

import "strconv"

func sol(a, b int) int {
	aWithB, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if aWithB >= 2 * a * b {
		return aWithB
	} else {
		return 2 * a * b
	}
}


func main() {
	sol(12, 3)
}