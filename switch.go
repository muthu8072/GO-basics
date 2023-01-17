package main

import "fmt"

func main() {
	/*fmt.Println("enter the number")
	var a int
	fmt.S10canln(&a)
	switch a {
	case 10:
		{
			fmt.Println("10 is right")
		}
	case 20:
		{
			fmt.Println("20 is right")
		}
	default:
		fmt.Println("not 10 and 20")

	}*/
	switch {
	case 10 > 5:
		fmt.Println("greater 10")
		fallthrough
	case 20 > 6:
		fmt.Println("greater 20")

	default:
		fmt.Println("Error")
	}
}
