package main

import "fmt"

func main() {
	fmt.Println("please fill the details")
lo:

	fmt.Println("Enter your age")
	var age int
	fmt.Scanln(&age)
	if age < 17 {
		goto lo
	} else {
		fmt.Println("you are eligible for vote")
	}
}
