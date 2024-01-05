package main

import (
	"fmt"
	"strings"
)

func main() {
    	var s1 string
    	fmt.Scan(&s1)
	var ans []string
	split := strings.Split(s1, "")

	for _, char := range split {
		if char >= "a" && char <= "z" {
			ans = append(ans, strings.ToUpper(char))
		} else if char >= "A" && char <= "Z" {
			ans = append(ans, strings.ToLower(char))
		}
	}

	fmt.Println(strings.Join(ans, ""))
}