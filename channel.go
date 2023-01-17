package main

import "fmt"

func main() {
	my := make(chan string)

	go func() {
		my <- "hello"
		my <- "muthu"
		my <- "what is this"
	}()

	text := <-my
	fmt.Println(text)
	fmt.Println(<-my)
	fmt.Println(<-my)

}
