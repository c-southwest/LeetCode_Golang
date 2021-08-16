package main

func main() {

}

func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		switch stack[len(stack)-1] {
		case '(':
			if v == ')' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		case '{':
			if v == '}' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		case '[':
			if v == ']' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
