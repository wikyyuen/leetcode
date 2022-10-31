package main

import "fmt"

//509. Fibonacci Number
//509. 斐波那契数

//var books []int
//
//func fib(n int) int {
//	books = make([]int, n)
//	return fn(n)
//}
//
//func fn(n int) int {
//	if n == 0 {
//		return 0
//	}
//	if n == 1 {
//		return 1
//	}
//	if books[n-1] == 0 {
//		books[n-1] = fn(n - 1)
//	}
//	if books[n-2] == 0 {
//		books[n-2] = fn(n - 2)
//	}
//	return books[n-1] + books[n-2]
//}

//solution 2

func fib(n int) int {
	if n < 2 {
		return n
	}
	cur := 0
	next := 1

	for i := 2; i < n; i++ {
		new := cur + next
		cur, next = next, new
	}
	return cur + next
}

func main() {
	n := 4

	result := fib(n)

	fmt.Println(result)
}
