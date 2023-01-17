package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number")
	fmt.Scanln(&a)
	if a >= 13 && a <= 18 {
		fmt.Println("you are teen age")
	} else {
		panic("not eligible")
	}
}
