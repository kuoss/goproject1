package main

import "fmt"

func main() {
	prepare("This should be a constant") // Noncompliant; 'This should ...' is duplicated 3 times
	execute("This should be a constant")
	release("This should be a constant")
}

func prepare(s string) {
	fmt.Println(s)
}
func execute(s string) {
	fmt.Println(s)
}
func release(s string) {
	fmt.Println(s)
}
