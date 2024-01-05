package main

import "fmt"

func solution(my_string string, overwrite_string string, s int) string {
	return my_string[:s] + overwrite_string + my_string[s+len(overwrite_string):]
    }

func main()  {
	fmt.Println(solution("He11oWor1d", "lloWorl", 2))
}