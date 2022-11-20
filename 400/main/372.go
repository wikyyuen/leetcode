package main

import (
	"fmt"
)

//372. Super Pow
//372. 超级次方

func myPow(a int, b int) int {
	ans := 1
	for b > 0 {
		if b%2 == 1 {
			ans = ans * a % 1337
		}
		a = a * a % 1337
		b /= 2
	}
	return ans
}

func superPow(a int, b []int) int {
	ans := 1
	for _, i := range b {
		ans = myPow(ans, 10) * myPow(a, i) % 1337
	}
	return ans
}

func main() {
	a := 2
	//b := []int{4, 3, 3, 8, 5, 2}
	//b := []int{1, 0}
	b := []int{3}

	res := superPow(a, b)

	fmt.Println(res)
}
