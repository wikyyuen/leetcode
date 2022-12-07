package main

import "fmt"

//921. Minimum Add to Make Parentheses Valid
//921. 使括号有效的最少添加

func minAddToMakeValid(s string) int {
	b := []byte(s)
	need := 0
	res := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '(' {
			need++
		} else {
			if need > 0 {
				need--
			} else {
				res++
				need = 0
			}
		}
	}
	return res + need
}

func main() {
	s := "())"
	res := minAddToMakeValid(s)
	fmt.Println(res)
}
