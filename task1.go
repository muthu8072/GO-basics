package main

import "fmt"

func main() {

	fmt.Println("Welcome please fill the all details")
	var a string
	fmt.Println("Enter your name")
	fmt.Scanln(&a)
	var b string
	fmt.Println("Enter your place")
	fmt.Scanln(&b)
	var c int
	fmt.Println("enter your date of birth")
	fmt.Scanln(&c)
	var d string
	fmt.Println("Enter your address")
	fmt.Scanln(&d)
	var e string
	fmt.Println("Enter your distric")
	fmt.Scanln(&e)
	var f float32
	fmt.Println("Enter your body weight")
	fmt.Scanln(&f)

	fmt.Println("your name is", a)
	fmt.Println("your place is", b)
	fmt.Println("your dob is", c)
	fmt.Println("your address is", d)
	fmt.Println("your distric is", e)
	fmt.Println("your bodyweight is", f)

	fmt.Println("please fill your voting detils")

	var adhar int32
	fmt.Println("Enter your aadhar card number")
	fmt.Scanln(&adhar)
	var fname string
	fmt.Println("Enter your father name")
	fmt.Scanln(&fname)
	var mname string
	fmt.Println("Enter your mother name")
	fmt.Scanln(&mname)
	var age int
	fmt.Println("Enter your age")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("your age is above 18,so your eligible for vote your full details")
		fmt.Println("your age is", age)
		fmt.Println("your adhar number is", adhar)
		fmt.Println("your father name is", fname)
		fmt.Println("your mother name is", mname)

	} else {
		fmt.Println("your not eligible for vote sorry!!!!")
	}
	fmt.Println("all details is correct press 0 show all details  press 1 your aadhar details press 2 all and adhar details")

	var mu int
	fmt.Println("plaese Enter the choose any one number")
	fmt.Scanln(&mu)
	//var val int
	//fmt.Scanln(&val)
	//var all int
	//fmt.Scanln(&all)

	if mu == 0 {

		fmt.Println("your name is", a)
		fmt.Println("your place is", b)
		fmt.Println("your dob is", c)
		fmt.Println("your address is", d)
		fmt.Println("your distric is", e)
		fmt.Println("your bodyweight is", f)

	} else if mu == 1 {

		fmt.Println("your age is", age)
		fmt.Println("your adhar number is", adhar)
		fmt.Println("your father name is", fname)
		fmt.Println("your mother name is", mname)
	} else if mu == 2 {
		fmt.Println("your name is", a)
		fmt.Println("your place is", b)
		fmt.Println("your dob is", c)
		fmt.Println("your address is", d)
		fmt.Println("your distric is", e)
		fmt.Println("your bodyweight is", f)
		fmt.Println("your age is", age)
		fmt.Println("your adhar number is", adhar)
		fmt.Println("your father name is", fname, "\n")
		fmt.Println("your mother name is", mname, "\n")

	} else {
		fmt.Println("Sorry you can choose in another number.....Thankyou for spend the time")
		fmt.Println("try again")
	}

}
