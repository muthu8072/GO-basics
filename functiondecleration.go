package main

import "fmt"

//
func mut() { //with parameter
	var a, b int
	a = 10
	b = 23
	c := a + b
	fmt.Println(c)
}

type aa struct {
	username string
	password int
}

func kali(a, b int) int {
	return a + b
}
func loop(c, d int) int {
	if c > d {
		fmt.Print("a is greater")
	} else if d > c {
		fmt.Println("d is greater")
	}
	return c + d
}

//func bb(aa)
func main() {

	mut()
	z := aa{username: "muthu", password: 234}
	fmt.Println(z)
	x := kali(0, 20)
	fmt.Println(x)
	l := loop(43, 20)
	fmt.Println(l)

}
