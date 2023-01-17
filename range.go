package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7}
	for j := range num {
		i := 2
		fmt.Println("slice item", j+i, "is", "number", num[j])
	}
}
