package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] == str[j] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(12321))
}
