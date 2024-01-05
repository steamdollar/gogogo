package main

func sol(a, b int, flag bool) int {
	if flag {
		return a + b
	} else {
		return a - b
	}
}

func main() {
	sol(1, 2, true)
}