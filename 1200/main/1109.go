package main

import "fmt"

//1109. Corporate Flight Bookings

//n^2
//func corpFlightBookings(bookings [][]int, n int) []int {
//	l := make([]int, n+1)
//	for _, item := range bookings {
//		for i := item[0]; i <= item[1]; i++ {
//			l[i] += item[2]
//		}
//	}
//	return l[1:]
//}

//2n
func corpFlightBookings(bookings [][]int, n int) []int {
	l := make([]int, n)
	for _, item := range bookings {
		l[item[0]-1] += item[2]
		if item[1] < n {
			l[item[1]] -= item[2]
		}
	}
	for i := 1; i < n; i++ {
		l[i] += l[i-1]
	}
	return l
}

func main() {
	booking := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	n := 5

	result := corpFlightBookings(booking, n)

	fmt.Println(result)
}
