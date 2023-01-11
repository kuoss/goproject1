package main

import "fmt"

func main() {
	fmt.Println("hello world1")
	fmt.Println("hello world2")
	fmt.Println("hello world3")
	fmt.Println("hello world4")
	fmt.Println("hello world5")
	fmt.Println("hello world")
}

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}
