package main

import (
	"fmt"
)

func sol(my_string string, index_list []int) {
	var ans string
	
	for i := 0; i < len(index_list); i++ {
		ans += string(my_string[index_list[i]])
	}
	fmt.Println(ans)
}

func main() {
	s1 := []int{16, 6, 5, 3, 12, 14, 11, 11, 17, 12, 7}
	sol("cvsgiorszzzmrpaqpe", s1)
}