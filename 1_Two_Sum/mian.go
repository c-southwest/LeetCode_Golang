package main

import "fmt"

func twoSum(nums []int, target int) []int {
	table := map[int]int{}

	for index, value := range nums {
		if _, ok := table[target-value]; ok == true {
			fmt.Println(index, table[target-value])
			return []int{index, table[target-value]}
		} else {
			table[value] = index
		}
	}

	return nil

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
	nums = []int{3, 2, 4}
	target = 6
	twoSum(nums, target)
}
