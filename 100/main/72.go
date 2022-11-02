package main

import "fmt"

//72. Edit Distance
//72. 编辑距离
var dpList [][]int

func minDistance(word1 string, word2 string) int {
	p1 := len([]byte(word1))
	p2 := len([]byte(word2))
	dpList = make([][]int, p1)
	for i := 0; i < p1; i++ {
		dpList[i] = make([]int, p2)
		for j := 0; j < p2; j++ {
			dpList[i][j] = -1
		}
	}

	return dp([]byte(word1), p1-1, []byte(word2), p2-1)
}

func dp(word1 []byte, p1 int, word2 []byte, p2 int) int {
	if p1 == -1 {
		return p2 + 1
	}
	if p2 == -1 {
		return p1 + 1
	}
	if dpList[p1][p2] != -1 {
		return dpList[p1][p2]
	}
	if word1[p1] == word2[p2] {
		return dp(word1, p1-1, word2, p2-1)
	}
	dpList[p1][p2] = min(
		dp(word1, p1, word2, p2-1)+1, //增加
		dp(word1, p1-1, word2, p2)+1, //删除
		dp(word1, p1-1, word2, p2-1)+1)
	return dpList[p1][p2]
}

func min(a int, b int, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}

func main() {
	word1 := "horse"
	word2 := "ros"

	res := minDistance(word1, word2)

	fmt.Println(res)

}
