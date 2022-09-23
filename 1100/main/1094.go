package main

import "fmt"

//1094. Car Pooling

func carPooling(trips [][]int, capacity int) bool {
	l := [1000]int{}
	for _, item := range trips {
		for i := item[1]; i < item[2]; i++ {
			l[i] += item[0]
		}
	}
	for _, i := range l {
		if i > capacity {
			return false
		}
	}
	return true

}

func main() {
	trips := [][]int{{2, 1, 5}, {3, 5, 7}}
	capacity := 4

	result := carPooling(trips, capacity)

	fmt.Println(result)
}
