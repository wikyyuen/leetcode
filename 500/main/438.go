package main

import (
	"fmt"
)

//438. Find All Anagrams in a String
//solution 1
//func findAnagrams(s string, p string) []int {
//	result := []int{}
//	length := len([]byte(p))
//	mp := make(map[string]int, length)
//	window := make(map[string]int, length)
//	for _, i2 := range p {
//		mp[string(i2)]++
//	}
//	left := 1
//	right := length
//	maxLen := len([]byte(s))
//	for i, i2 := range s {
//		if i >= length {
//			break
//		}
//		window[string(i2)]++
//	}
//	if reflect.DeepEqual(window, mp) {
//		result = append(result, 0)
//	}
//	for right < maxLen {
//		window[string(s[right])]++
//		window[string(s[left-1])]--
//		if window[string(s[left-1])] == 0 {
//			delete(window, string(s[left-1]))
//		}
//		if reflect.DeepEqual(window, mp) {
//			result = append(result, left)
//		}
//		left++
//		right++
//	}
//	return result
//}

//solution2
func findAnagrams(s string, p string) []int {
	counter := make([]int, 26)
	for i := 0; i < len(p); i++ {
		counter[p[i]-97]--
	}
	result := []int{}
	for i := 0; i < len(s); i++ {
		counter[s[i]-97]++
		if i < len(p) {
			if checkAllZero(counter) {
				result = append(result, i-len(p)+1)
			}
			continue
		}
		counter[s[i-len(p)]-97]--
		if checkAllZero(counter) {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

func checkAllZero(counter []int) bool {
	for _, i := range counter {
		if i != 0 {
			return false
		}
	}
	return true
}

func main() {
	//s := "cbaebabacd"
	//p := "abc"
	s := "abab"
	p := "ab"

	result := findAnagrams(s, p)

	fmt.Println(result)
}
