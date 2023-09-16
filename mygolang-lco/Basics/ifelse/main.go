package main

import "fmt"

func main() {

	fmt.Println("If else golang")

	loginCount := 23

	if loginCount < 10{
		fmt.Println("Less than 10")
	} else if loginCount > 10{
		fmt.Println("Greater than 10")
	} else{
		fmt.Println("Equals 10")
	}

	if num:= 4; num%2 == 0{
		fmt.Println("Number is even")
	} else{
		fmt.Println("Number is odd")
	}

	// way to check err is nil or not
	// if err != nil{
	// 	fmt.Println("LoginCount exist")
	// }

}
