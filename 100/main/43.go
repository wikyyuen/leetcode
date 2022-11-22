package main

import (
	"fmt"
	"strconv"
)

//43. Multiply Strings
//43. 字符串相乘

func multiply(num1 string, num2 string) string {
	n1 := []byte(num1)
	n2 := []byte(num2)
	l1 := len(num1)
	l2 := len(num2)
	res := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			n1i, _ := strconv.Atoi(string(n1[i]))
			n2i, _ := strconv.Atoi(string(n2[j]))
			mul := n1i * n2i
			p1 := i + j
			p2 := p1 + 1 //进位
			sum := res[p2] + mul
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}
	str := ""
	for i < len(res) {
		str += strconv.Itoa(res[i])
		i++
	}
	if len(str) == 0 {
		return "0"
	}
	return str
}

func main() {
	num1 := "123"
	num2 := "456"

	res := multiply(num1, num2)

	fmt.Println(res)
}
