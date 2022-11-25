package main

import (
	"fmt"
	"strconv"
)

//241. Different Ways to Add Parentheses
//241. 为运算表达式设计优先级

func diffWaysToCompute(expression string) []int {
	res := []int{}
	for i := 0; i < len(expression); i++ {
		if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					if expression[i] == '-' {
						res = append(res, l-r)
					} else if expression[i] == '+' {
						res = append(res, l+r)
					} else if expression[i] == '*' {
						res = append(res, l*r)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		temp, _ := strconv.Atoi(expression)
		return []int{temp}
	}
	return res
}

func main() {
	expression := "2*3-4*5"

	res := diffWaysToCompute(expression)

	fmt.Println(res)
}
