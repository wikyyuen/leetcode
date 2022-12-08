package main

import "fmt"

//22. Generate Parentheses
//22. 括号生成
var resList []string

func generateParenthesis(n int) []string {
	resList = []string{}
	traverse("(", 1, n-1, n)
	return resList
}

func traverse(s string, need int, left int, right int) {
	if right == 0 && left == 0 {
		resList = append(resList, s)
	}
	if need > 0 && right > 0 {
		traverse(s+")", need-1, left, right-1)
	}
	if left > 0 {
		traverse(s+"(", need+1, left-1, right)
	}
}

func main() {
	n := 3
	res := generateParenthesis(n)

	fmt.Println(res)
}
