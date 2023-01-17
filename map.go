package main

import "fmt"

func main() {
	a := map[int]float64{2: 20.4, 10: 30.5, 23: 30.5}
	a[67] = 20.7 //update
	//a["mmm"] = 30
	fmt.Println(a)
	delete(a, 2)
	ru := make(map[string]int)
	ru["manish"] = 10

	v, o := ru["manish"]
	fmt.Printf(v, o)

	fmt.Println(a)
}
