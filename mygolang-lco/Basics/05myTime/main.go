package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// we use 01-02-2006 formatting date
	// 01-02-2006 Monday formatting date and day
	// 01-02-2006 15:04:05 Monday formatting date time and day



	createdDate := time.Date(2002, time.March, 20, 9, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}