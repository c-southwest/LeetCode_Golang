package main

import "fmt"

func main() {
	s := "bb"
	r := longestPalindrome3(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	start := 0
	end := 0
	//外层循环
	for i := 0; i < len(s); i++ {
		str := s[i:len(s)]
		//每次循环end-1
	FLAG:
		for j := len(str) - 1; j > 0; j-- {
			if (end - start) > (j) {
				break
			}
			//验证所选范围是否满足要求
			for first, last := 0, j; first < last; first, last = first+1, last-1 {
				if str[first] == str[last] {
					continue
				} else {
					//不满足条件
					continue FLAG
				}
			}
			//满足条件
			if (end - start) < (j) {
				start = i
				end = i + j
			}
		}
	}
	return string(s[start : end+1])
}

func longestPalindrome2(s string) string {
	start := 0
	end := 0
	//外层循环
	for i := 0; i < len(s); i++ {
	FLAG:
		for j := len(s) - 1; j > i; j-- {
			if (end - start) > (j - i) {
				break
			}
			//验证所选范围是否满足要求
			for first, last := i, j; first < last; first, last = first+1, last-1 {
				if s[first] == s[last] {
					continue
				} else {
					//不满足条件
					continue FLAG
				}
			}
			//满足条件
			if (end - start) < (j - i) {
				start = i
				end = j
			}
		}
	}
	return string(s[start : end+1])
}

//中心扩散法
func longestPalindrome3(s string) string {
	start := 0
	end := 0
	//外层循环
	for i := 0; i < len(s); i++ {
		//开始扩散
		//奇数
		for left, right := i, i; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] == s[right] {
				if (end - start) < (right - left) {
					start = left
					end = right
				}
			} else {
				break
			}
		}

		//偶数
		for left, right := i, i+1; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] == s[right] {
				if (end - start) < (right - left) {
					start = left
					end = right
				}
			} else {
				break
			}
		}
	}
	return s[start : end+1]
}
