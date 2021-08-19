package main

func main() {

}

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v == val {

		} else {
			nums[index] = v
			index++
		}
	}
	return index
}
