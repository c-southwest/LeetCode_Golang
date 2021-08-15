package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world  "
	fmt.Println(strings.Trim(s, " "))
}
