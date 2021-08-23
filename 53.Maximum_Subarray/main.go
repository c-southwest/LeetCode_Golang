package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(nums)
	fmt.Println(r)
	nums2 := []int{-1}
	r2 := maxSubArray(nums2)
	fmt.Println(r2)
}

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, i := range nums {
		sum += i
		if sum > max {
			max = sum
		}
		if sum <= 0 {
			sum = 0
			continue
		}
	}
	return max
}
