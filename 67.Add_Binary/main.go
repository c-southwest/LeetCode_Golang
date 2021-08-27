package main

import "fmt"

func main() {
	a := "101"
	b := "10101"
	r := addBinary(a, b)
	fmt.Println(r)
}

func addBinary(a string, b string) string {
	var result []rune
	if len(a) >= len(b) {
		result = []rune(a)
		flag := 0
		i := 0
		for ; i <= len(b)-1; i++ {
			if b[len(b)-1-i] == '1' && result[len(a)-1-i] == '1' {
				if flag == 1 {
					flag = 1
				} else {
					result[len(a)-1-i] = '0'
					flag = 1
				}
				continue
			} else if b[len(b)-1-i] == '1' && result[len(a)-1-i] == '0' {
				if flag == 1 {

				} else {
					result[len(a)-1-i] = '1'
					flag = 0
				}
				continue

			} else if b[len(b)-1-i] == '0' && result[len(a)-1-i] == '1' {
				if flag == 1 {
					result[len(a)-1-i] = '0'
				} else {
					result[len(a)-1-i] = '1'
					flag = 0
				}
				continue

			} else if b[len(b)-1-i] == '0' && result[len(a)-1-i] == '0' {
				if flag == 1 {
					flag = 0
					result[len(a)-1-i] = '1'
				} else {
					result[len(a)-1-i] = '0'
					flag = 0
				}
				continue
			}
		}
		if flag == 1 {
			for i < len(a) {
				if result[len(a)-1-i] == '1' {
					result[len(a)-1-i] = '0'
					i++
				} else {
					if flag == 1 {
						result[len(a)-1-i] = '1'
					}
					return string(result)
				}
			}
			if flag == 1 && i == len(a) {
				//result = append([]rune{'1'}, result...)
				return "1" + string(result)
			}

		}
	} else {
		result = []rune(b)
		flag := 0
		i := 0
		for ; i <= len(a)-1; i++ {
			if a[len(a)-1-i] == '1' && result[len(b)-1-i] == '1' {
				if flag == 1 {
					flag = 1
				} else {
					result[len(b)-1-i] = '0'
					flag = 1
				}
				continue
			} else if a[len(a)-1-i] == '1' && result[len(b)-1-i] == '0' {
				if flag == 1 {

				} else {
					result[len(b)-1-i] = '1'
					flag = 0
				}
				continue

			} else if a[len(a)-1-i] == '0' && result[len(b)-1-i] == '1' {
				if flag == 1 {
					result[len(b)-1-i] = '0'
				} else {
					result[len(b)-1-i] = '1'
					flag = 0
				}
				continue

			} else if a[len(a)-1-i] == '0' && result[len(b)-1-i] == '0' {
				if flag == 1 {
					result[len(b)-1-i] = '1'
					flag = 0
				} else {
					result[len(b)-1-i] = '0'
					flag = 0
				}
				continue
			}
		}
		if flag == 1 {
			for i < len(b) {
				if result[len(b)-1-i] == '1' {
					result[len(b)-1-i] = '0'
					i++
				} else {
					if flag == 1 {
						result[len(b)-1-i] = '1'
					}
					return string(result)
				}
			}
			if flag == 1 && i == len(b) {
				//result = append([]rune{'1'}, result...)
				return "1" + string(result)
			}

		}
	}
	return string(result)
}
