package main

import "fmt"

//51. N-Queens
//51. N 皇后
var length int
var correct [][]string

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	length = n
	correct = [][]string{}
	for i := 0; i < length; i++ {
		cur := []int{i}
		traverse3(cur, 1)
	}
	return correct
}

func traverse3(cur []int, index int) {
	for i := 0; i < length; i++ {
		if isValid(cur, i, index) {
			temp := make([]int, len(cur))
			copy(temp, cur)
			temp = append(temp, i)
			if index == length-1 {
				correct = append(correct, transform(temp))
				return
			}
			traverse3(temp, index+1)
		}
	}
}

func isValid(cur []int, x int, y int) bool {
	for index, item := range cur {
		if !isValidSingle(item, index, x, y) {
			return false
		}
	}
	return true
}

func transform(result []int) []string {
	res := []string{}
	for _, i2 := range result {
		temp := []byte("")
		for i := 0; i < length; i++ {
			if i == i2 {
				temp = append(temp, 'Q')
			} else {
				temp = append(temp, '.')
			}
		}
		res = append(res, string(temp))
	}
	return res
}

func isValidSingle(x1 int, y1 int, x2 int, y2 int) bool {
	if x1 == x2 || y1 == y2 {
		return false
	}
	if (x1-x2) == (y1-y2) || (x1-x2) == (y2-y1) {
		return false
	}
	return true
}

func main() {
	n := 4

	res := solveNQueens(n)

	fmt.Println(res)

	//fmt.Println(isValidSingle(2, 2, 3, 1))
}
