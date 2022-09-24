package main

import "fmt"

//151. Reverse Words in a String
//solution 1
//func reverseWords(s string) string {
//	last_type := true
//	newByte := []byte(s)
//	var byteArray []byte
//	newString := ""
//	len := len(newByte)
//	for i := 0; i < len; i++ {
//		if newByte[i] != ' ' {
//			if last_type == true && byteArray != nil {
//				newString = " " + string(byteArray) + newString
//				byteArray = []byte("")
//			}
//			byteArray = append(byteArray, newByte[i])
//			last_type = false
//		} else {
//			last_type = true
//		}
//	}
//	newString = string(byteArray) + newString
//	return newString
//}

//solution 2
func reverseWords(s string) string {
	lastType := true
	newByte := []byte(s)
	newString := ""
	length := len(newByte)
	lastStart := 0
	lastEnd := 0
	for i := 0; i < length; i++ {
		if newByte[i] != ' ' {
			if lastType == true {
				if lastEnd != 0 {
					newString = " " + string(newByte[lastStart:lastEnd]) + newString
					lastEnd = 0
				}
				lastStart = i
			}
			lastType = false
		} else {
			if lastType == false {
				lastEnd = i
			}
			lastType = true
		}
	}
	if lastEnd == 0 {
		lastEnd = length
	}
	newString = string(newByte[lastStart:lastEnd]) + newString
	return newString
}

func main() {
	//s := "the sky is blue"
	s := "  hello world  "
	result := reverseWords(s)
	fmt.Println(result)
}
