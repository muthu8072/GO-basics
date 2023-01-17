package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) fullname() {
	fmt.Println(p.fname, "", p.lname)

}
func main() {
	a := person{"krish", "raja"}
	a.fullname()

}
