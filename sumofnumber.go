package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	sum := 0
	var total int
	for _, k := range arr {
		sum += k
		total = sum

	}
	fmt.Println(total)

}
