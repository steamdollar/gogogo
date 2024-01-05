package main

import (
	"fmt"
	"strings"
)

func sol(str string) {
	var ans string
	for i := 0; i <len(str); i++ {
		ans += strings.ToLower(string(str[i]))
	}
	fmt.Println(ans)
}

func main() {
    sol("QIOWknasd")
}