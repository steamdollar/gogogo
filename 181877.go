package main

import (
	"strings"
)

func sol(str string)  {
	var ans string
	for _, v := range str {
		ans += strings.ToUpper(string(v))
	}
	return ans
}

func main() {
    sol("qwAS")
}