package main

import (
	"fmt"
	"time"
)

func main() {
	currenttime := time.Now()
	currentday := time.Now().Weekday()
	fmt.Println("current time is", currenttime.String())
	fmt.Println("today day is", currentday)
	//fmt.Println("YYYY-MM-DD : ", currenttime.Format("2017-09-07"))
}
