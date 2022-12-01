package main

import (
	"fmt"
)

//752. Open the Lock
//752. 打开转盘锁
//BFS
func openLock(deadends []string, target string) int {
	deadMap := map[string]bool{}
	visitedMap := map[string]bool{}
	step := 0
	for i := 0; i < len(deadends); i++ {
		deadMap[deadends[i]] = true
	}
	res := []string{}
	res = append(res, "0000")
	visitedMap["0000"] = true

	for len(res) > 0 {
		n := len(res)
		for i := 0; i < n; i++ {
			cur := res[0]
			res = res[1:]
			if _, ok := deadMap[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			for j := 0; j < 4; j++ {
				//每一位处理两次
				next := upOne(cur, j)
				if _, ok := visitedMap[next]; !ok {
					visitedMap[next] = true
					res = append(res, next)
				}
				next = downOne(cur, j)
				if _, ok := visitedMap[next]; !ok {
					visitedMap[next] = true
					res = append(res, next)
				}
			}
		}
		step++
	}
	return -1
}

func upOne(str string, index int) string {
	b := []byte(str)
	if b[index] == '9' {
		b[index] = '0'
	} else {
		b[index] += 1
	}
	return string(b)
}

func downOne(str string, index int) string {
	b := []byte(str)
	if b[index] == '0' {
		b[index] = '9'
	} else {
		b[index] -= 1
	}
	return string(b)
}

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"

	res := openLock(deadends, target)

	fmt.Println(res)
}
