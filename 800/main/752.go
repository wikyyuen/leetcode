package main

import (
	"fmt"
)

//752. Open the Lock
//752. 打开转盘锁
//双向BFS
func openLock(deadends []string, target string) int {
	deadMap := map[string]bool{}
	visitedMap := map[string]bool{}
	backVisitedMap := map[string]bool{}
	step := 0
	for i := 0; i < len(deadends); i++ {
		deadMap[deadends[i]] = true
	}
	forward := []string{}
	forward = append(forward, "0000")
	visitedMap["0000"] = true

	backward := []string{}
	backward = append(backward, target)
	backVisitedMap[target] = true

	for len(forward) > 0 || len(backward) > 0 {
		n := len(forward)
		for i := 0; i < n; i++ {
			cur := forward[0]
			forward = forward[1:]
			if _, ok := deadMap[cur]; ok {
				continue
			}
			if _, ok := backVisitedMap[cur]; ok {
				return step
			}
			for j := 0; j < 4; j++ {
				//每一位处理两次
				next := upOne(cur, j)
				if _, ok := visitedMap[next]; !ok {
					visitedMap[next] = true
					forward = append(forward, next)
				}
				next = downOne(cur, j)
				if _, ok := visitedMap[next]; !ok {
					visitedMap[next] = true
					forward = append(forward, next)
				}
			}
		}
		step++
		n2 := len(backward)
		for i := 0; i < n2; i++ {
			curBack := backward[0]
			backward = backward[1:]
			if _, ok := deadMap[curBack]; ok {
				continue
			}
			if _, ok := visitedMap[curBack]; ok {
				return step
			}
			for j := 0; j < 4; j++ {
				//每一位处理两次
				next := upOne(curBack, j)
				if _, ok := backVisitedMap[next]; !ok {
					backVisitedMap[next] = true
					backward = append(backward, next)
				}
				next = downOne(curBack, j)
				if _, ok := backVisitedMap[next]; !ok {
					backVisitedMap[next] = true
					backward = append(backward, next)
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
