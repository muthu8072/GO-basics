package main

import "fmt"

func main() {
	for a := 0; a < 5; a++ {
		if a == 3 {
			continue
		}
		fmt.Println(a)
	}
}
