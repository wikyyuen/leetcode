package main

import "fmt"

//14. Longest Common Prefix

func longestCommonPrefix(strs []string) string {
	a := []byte(strs[0])
	strs = strs[1:]
	same := true
	newStr := []byte{}
	for i := 0; i < len(a); i++ {
		same = true
		for _, str := range strs {
			if len([]byte(str)) < i+1 || a[i] != []byte(str)[i] {
				same = false
				break
			}
		}
		if same {
			newStr = append(newStr, a[i])
		} else {
			break
		}
	}
	return string(newStr)
}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"ab", "a"}

	result := longestCommonPrefix(strs)

	fmt.Println(result)
}
