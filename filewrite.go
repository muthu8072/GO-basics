package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.txt") //create file
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("hiii muthuraja") //write file
	file.Close()                       //close file

	stream, err := ioutil.ReadFile("file.txt") //read file
	if err != nil {
		log.Fatal(err)
	}
	result := string(stream)
	fmt.Println(result)

}
