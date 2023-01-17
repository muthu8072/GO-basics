package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 4, 7, 9, 5, 10}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)
}
