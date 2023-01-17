package main

import "fmt"

func main() {

	var a int
	fmt.Printf("Enter the number")
	fmt.Scanln(&a)

	if a == 10 {
		fmt.Print("even number")
	} else {
		fmt.Print("add")
	}
}
