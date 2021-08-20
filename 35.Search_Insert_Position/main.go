package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	r := searchInsert(nums, 2)
	fmt.Println("r:", r)
}

func searchInsert(nums []int, target int) int {
	a, b := 0, len(nums)-1
	return do(nums, a, b, target)
}

func do(nums []int, a, b, target int) int {
	if a >= b {
		if target <= nums[a] {
			return a
		} else {
			return a + 1
		}
	}
	mid := (a + b) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < target {
		return do(nums, mid+1, b, target)
	} else {
		return do(nums, a, mid-1, target)
	}
}
