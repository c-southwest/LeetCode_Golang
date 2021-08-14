package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	r := longestCommonPrefix(strs)
	fmt.Println(r)
}

func longestCommonPrefix(strs []string) string {
	s := strs[0]
	end := len(strs[0])
	for _, str := range strs {
		if len(str) < end {
			end = len(str)
		}
		for i := 0; i < end; i++ {
			if str[i] == s[i] {
				continue
			} else {
				end = i
				break
			}
		}
	}
	return string(s[0:end])
}
