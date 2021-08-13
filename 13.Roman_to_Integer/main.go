package main

import "fmt"

func main() {
	r := romanToInt("MCMXCIV")
	fmt.Println(r)
}

func romanToInt(s string) int {
	table := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	sum := 0
	for k, v := range s {
		if k == len(s)-1 {
			sum += table[string(v)]
			break
		}
		if table[string(v)] >= table[string(s[k+1])] {
			sum += table[string(v)]
		} else {
			sum -= table[string(v)]
		}
	}
	return sum
}
