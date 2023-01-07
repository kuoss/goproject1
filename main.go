package main

import "fmt"

func main() {
        fmt.Println("hello world")
	numbers := []int{1, 2, 3, 4}
	if len(numbers) < 0 {
		fmt.Println("unreachable code")
	}
}
