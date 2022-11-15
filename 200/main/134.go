package main

import "fmt"

//134. Gas Station

func canCompleteCircuit(gas []int, cost []int) int {
	sum, minSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			minSum = sum
			start = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	if start == len(gas) {
		return 0
	}
	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	res := canCompleteCircuit(gas, cost)

	fmt.Println(res)
}
