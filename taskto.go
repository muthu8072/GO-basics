package main

import "fmt"

type muthu struct {
	name    string
	age     int
	place   string
	distric string
}

func cal() {
	var a, b int
	a = 10
	b = 23
	c := a + b
	fmt.Println(c)
}
func main() {
	x := muthu{name: "muthuraja", age: 21, place: "kallidai", distric: "Tirunelveli"}
	fmt.Println(x)
	cal()
}
