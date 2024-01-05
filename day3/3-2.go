package main

import "fmt"


func sol(arr []string) {
	
	w := make([]byte, len(arr))
	
	for i, alphabet := range arr {
		fmt.Println(alphabet)
		w[i] = alphabet[0]
	}
	
	
	fmt.Println(string(w))
	
	//
	var ans string
	for i:=0; i < len(arr); i++ {
		ans += arr[i]
	}
	fmt.Println(ans)
}

func main() {
    s1 := []string{"a", "b", "c"}
    
    sol(s1)
}