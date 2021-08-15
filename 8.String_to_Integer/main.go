package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	r := myAtoi("   -42")
	fmt.Println(r)

}

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	number := big.NewInt(0)
	if len(s) == 0 {
		return 0
	}
	isPos := true
	start := 0
	if string(s[0]) == "+" {
		isPos = true
		start = 1
	} else if string(s[0]) == "-" {
		isPos = false
		start = 1
	}
	for _, digit := range s[start:] {
		if digit >= '0' && digit <= '9' {
			number = number.Mul(number, big.NewInt(10))
			number = number.Add(number, big.NewInt(int64(digit-'0')))
		} else {
			break
		}
	}
	if isPos {

	} else {
		number = number.Neg(number)
	}
	if number.Cmp(big.NewInt(math.MaxInt32)) <= 0 {
		if number.Cmp(big.NewInt(math.MinInt32)) >= 0 {
			str := number.String()
			n, _ := strconv.Atoi(str)
			return n
		} else {
			return math.MinInt32
		}

	} else {
		return math.MaxInt32
	}
}
