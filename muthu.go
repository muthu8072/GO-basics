package main

import "fmt"

type data struct {
	name   string
	age    int
	place  string
	weight float32
}

type home struct {
	data
	address string
	stnum   int
	dis     string
}

func main() {
	y := home{data: data{name: "kali", age: 20, place: "ambai"}, address: "nadu", stnum: 2, dis: "Tirunelveli"}
	x := data{name: "Muthuraja", age: 21, place: "Kallidai", weight: 73.5}
	fmt.Printf("%v", x)
	fmt.Print("your age is")
	fmt.Printf("%v", y.data.age)

}
