package main

import "fmt"

//392. Is Subsequence
//392. 判断子序列

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	sB := []byte(s)
	tB := []byte(t)
	sN := len(sB)
	tN := len(tB)
	j := 0
	for i := 0; i < tN; i++ {
		if sB[j] == tB[i] {
			j++
			if j == sN {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"

	result := isSubsequence(s, t)

	fmt.Println(result)
}
