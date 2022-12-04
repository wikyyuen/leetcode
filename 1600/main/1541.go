package main

//1541. Minimum Insertions to Balance a Parentheses String
//1541. 平衡括号字符串的最少插入次数

func minInsertions(s string) int {
	res := 0
	need := 0
	for _, r := range s {
		if r == '(' {
			need += 2
			if need%2 == 1 {
				need--
				res++
			}
		} else {
			if need > 0 {
				need--

			} else {
				res++
				need = 1
			}
		}
	}
	return res + need
}
