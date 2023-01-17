package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter the number")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)
}
