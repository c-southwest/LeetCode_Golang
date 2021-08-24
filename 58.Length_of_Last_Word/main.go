package main

import "fmt"

func main() {
	s := "Hello World"
	r := lengthOfLastWord(s)
	fmt.Println(r)
}

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	length := 0
	for index >= 0 {
		if s[index] == ' ' {
			index--
			continue
		}
		break
	}
	for index >= 0 {
		if s[index] == ' ' {
			return length
		}
		length += 1
		index--
	}
	return length
}
