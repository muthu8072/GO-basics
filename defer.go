package main

import "fmt"

func main() {
	muthu("raja")
	defer muthu("first")
	defer muthu("first")
	defer muthu("second")
	defer muthu("third ")
}

func muthu(a string) {
	fmt.Println(a)
}
