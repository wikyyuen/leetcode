package main

import "fmt"

//20. Valid Parentheses
//20. 有效的括号

func isValid(s string) bool {
	stack := []byte{}
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if b[i] == '(' || b[i] == '[' || b[i] == '{' {
			stack = append(stack, b[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if b[i] == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if b[i] == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if b[i] == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	s := "()[]{}"

	res := isValid(s)

	fmt.Println(res)
}
