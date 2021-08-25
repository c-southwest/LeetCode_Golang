package main

func main() {

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
			if i == 0 {
				x := append([]int{1}, digits...)
				return x
			}
			continue
		}
	}
	return digits
}
