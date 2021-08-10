package main

import "fmt"

func main() {
	s := "hello world"
	for k, v := range s {
		fmt.Println(k, v)
	}
}
