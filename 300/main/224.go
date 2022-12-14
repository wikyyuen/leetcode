package main

import (
	"fmt"
)

//224. Basic Calculator
//224. 基本计算器

func calculate(s string) int {
	s += "+0"
	n := len(s)
	caList := []int{}
	cur := 0
	tempCur := 0
	curNums := []byte{}
	var curSymbol byte
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' {
			need := 1
			for j := i + 1; j < n; j++ {
				if s[j] == '(' {
					need++
					continue
				}
				if s[j] == ')' {
					need--
					if need == 0 {
						tempCur = calculate(s[i+1 : j])
						i = j
						break
					}
				}
			}
			continue
		}
		if s[i] == '-' || s[i] == '+' || s[i] == '*' || s[i] == '/' {
			//转换当前数字
			if len(curNums) > 0 {
				tempCur = 0
				curNumLength := len(curNums)
				for k := 0; k < curNumLength; k++ {
					tempCur += (int(curNums[k]) - int('0')) * myPow(curNumLength-1-k)
				}
				curNums = []byte{}
			}
			if curSymbol == '/' {
				cur = cur / tempCur
			} else if curSymbol == '*' {
				cur = cur * tempCur
			} else if curSymbol == '-' {
				cur -= tempCur
			} else {
				cur += tempCur
			}
			if s[i] == '-' || s[i] == '+' {
				caList = append(caList, cur)
				cur = 0
			}
			curSymbol = s[i]
		} else {
			curNums = append(curNums, s[i])
		}
	}
	cur = 0
	for i := 0; i < len(caList); i++ {
		cur += caList[i]
	}
	return cur
}

func myPow(n int) int {
	temp := 1
	for i := 0; i < n; i++ {
		temp *= 10
	}
	return temp
}

func main() {
	//s := "1 + 1"
	//s := " 2-1 + 2 "
	s := "(1+(4+5+2)-3)+(6+8)"

	result := calculate(s)

	fmt.Println(result)
}
