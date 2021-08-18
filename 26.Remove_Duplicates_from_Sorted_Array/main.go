package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r := removeDuplicates2(nums)
	fmt.Println(r)
	fmt.Println(nums[:r])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	result := 1
	flag := nums[0]
	for k := 1; k < len(nums); k++ {
		if nums[k] == flag {
			i := k + 1
			for ; i < len(nums); i++ {
				if nums[i] != nums[k] {
					for j := k; j < i; j++ {
						nums[j] = nums[i]
					}
					flag = nums[i]
					result += 1
					break
				}

			}
			if i == len(nums) {
				return result
			}

		} else {
			result += 1
			flag = nums[k]
		}
	}
	return result
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for _, v := range nums[1:] {
		if v != nums[index] {
			index += 1
			nums[index] = v
		}
	}
	return index + 1
}
