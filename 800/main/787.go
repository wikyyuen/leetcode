package main

import (
	"fmt"
	"strconv"
)

//787. Cheapest Flights Within K Stops
//787. K 站中转内最便宜的航班

var dpList map[string]int

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	flightsGraph := make([][][]int, n)
	for i := 0; i < len(flights); i++ {
		flightsGraph[flights[i][0]] = append(flightsGraph[flights[i][0]], []int{flights[i][1], flights[i][2]})
	}
	dpList = make(map[string]int)
	return dp(flightsGraph, src, dst, k)
}

func dp(flightsGraph [][][]int, src int, dst int, k int) int {
	if k == -1 {
		return -1
	}
	if value, ok := dpList[strconv.Itoa(src)+","+strconv.Itoa(k)]; ok {
		return value
	}
	max := 9999999
	for i := 0; i < len(flightsGraph[src]); i++ {
		var temp int
		if flightsGraph[src][i][0] == dst {
			temp = 0
		} else {
			temp = dp(flightsGraph, flightsGraph[src][i][0], dst, k-1)
		}
		if temp == -1 {
			continue
		}
		temp += flightsGraph[src][i][1]
		if temp < max {
			max = temp
		}
	}
	if max == 9999999 {
		max = -1
	}
	dpList[strconv.Itoa(src)+","+strconv.Itoa(k)] = max
	return max
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
