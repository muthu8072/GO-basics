package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type muthu struct {
		name string
		age  int
		cars []string
	}

	p := &muthu{name: "ram", age: 23, cars: []string{"tata", "maruthu", "hundai"}}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}
