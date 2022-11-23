package main

import (
	"fmt"
)

//787. Cheapest Flights Within K Stops
//787. K 站中转内最便宜的航班

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	minDst := make([]int, n)
	for i := 0; i < n; i++ {
		minDst[i] = 999999
	}
	minDst[src] = 0
	for i := 0; i < k+1; i++ {
		temp := make([]int, n)
		copy(temp, minDst)
		for _, flight := range flights {
			s, d, l := flight[0], flight[1], flight[2]
			if minDst[s] != 999999 && minDst[s]+l < temp[d] {
				temp[d] = minDst[s] + l
			}
		}
		copy(minDst, temp)
	}
	if minDst[dst] == 999999 {
		return -1
	}
	return minDst[dst]
}

func main() {
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 3
	k := 1

	res := findCheapestPrice(n, flights, src, dst, k)

	fmt.Println(res)
}
