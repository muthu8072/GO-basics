package main

import "fmt"

type stdetails struct {
	name, address, distric, place string
}

func main() {

	au := make(map[string]stdetails)

	au["muthuraja"] = stdetails{name: "muthuraja", address: "karanthai", distric: "tvl", place: "kallidai"}
	au["kalimuthu"] = stdetails{name: "kali", address: "mmmm"}
	fmt.Println(au)

}
