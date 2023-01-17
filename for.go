package main

import "fmt"

func main() { //nested for loop
	//innerLoopControl := 1
	for a := 1; a < 5; a++ {
		fmt.Println(a)
		for b := 1; b <= a+1; b++ {
			//fmt.Println(" ")
			fmt.Println(b)
		}

	}
}
