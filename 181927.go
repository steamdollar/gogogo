package main

import "fmt"

func sol(arr []int ) []int {
	len := len(arr)
	if arr[len -1] > arr[len-2] {
		arr = append(arr, arr[len-1] - arr[len-2])
	} else {
		arr = append(arr, arr[len-1] * 2)
	}
	return arr
}

func main() {
	arr := []int{2, 1, 6}
    fmt.Println(sol(arr))
}