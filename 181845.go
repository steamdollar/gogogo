package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func sol(n int) string {
	fmt.Println(reflect.TypeOf(n))
	return strconv.Itoa(n)
}

func main() {
    fmt.Println(sol(123))
    
}