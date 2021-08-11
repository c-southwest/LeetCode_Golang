package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = x * -1
	}
	str := fmt.Sprintf("%d", x)
	fmt.Println(str)
	var slc []int
	for i := range str {
		num, _ := strconv.Atoi(string(str[i]))
		slc = append(slc, num)
	}
	fmt.Println(slc)
	for i := 0; i < len(slc)/2; i++ {
		slc[i], slc[len(str)-1-i] = slc[len(str)-1-i], slc[i]
	}
	fmt.Println(slc)

	result := ""
	leading := true
	for i := range slc {
		if leading == true && slc[i] == 0 {
			continue
		} else {
			leading = false
			result = result + strconv.Itoa(slc[i])
		}
	}
	if neg == true {
		result = "-" + result
	}

	fmt.Println(result)
	result_int, _ := strconv.Atoi(result)
	if result_int > math.MaxInt32 || result_int < math.MinInt32 {
		return 0
	}
	return result_int
}

func main() {
	r := reverse2(-123456)
	fmt.Println(r)
	fmt.Println(int(math.Pow(2, 31)))
	fmt.Println(math.MaxInt32)
}

func reverse2(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = x * -1
	}
	slc := []rune(strconv.Itoa(x))
	fmt.Println(slc)

	for i := 0; i < len(slc)/2; i++ {
		slc[i], slc[len(slc)-1-i] = slc[len(slc)-1-i], slc[i]
	}
	fmt.Println(slc)
	result := string(slc)
	result_int, _ := strconv.Atoi(result)
	if neg {
		result_int *= -1
	}

	if result_int > math.MaxInt32 || result_int < math.MinInt32 {
		return 0
	}
	return result_int
}
