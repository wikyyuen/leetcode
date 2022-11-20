package main

import (
	"fmt"
)

//372. Super Pow
//372. 超级次方
var gBase = 1337

func myPow(a int, k int) int {
	if k == 0 {
		return 1
	}
	a %= gBase

	if k%2 == 1 {
		return (a * myPow(a, k-1)) % gBase
	} else {
		sub := myPow(a, k/2)
		return (sub * sub) % gBase
	}
}

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := myPow(a, last)
	part2 := myPow(superPow(a, b), 10)
	return part1 * part2 % gBase
}

func main() {
	a := 1
	b := []int{4, 3, 3, 8, 5, 2}
	//b := []int{1, 0}

	res := superPow(a, b)

	fmt.Println(res)
}
