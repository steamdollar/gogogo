package main

import (
	"fmt"
)


func sol(str1 string, str2 string) {
	
	var ans string
	fmt.Println(len(str1))
	for i :=0; i < len(str1); i++ {
		ans += str1[i:i+1]
		ans += str2[i:i+1]
	}
	
	fmt.Println(ans)
}

func main() {
    s1 := "aaaaa"
    s2 := "bbbbb"
    
    sol(s1, s2)
}