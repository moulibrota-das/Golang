package main

import "fmt"

func main() {
	defer fmt.Println("Defer one") 
	defer fmt.Println("Defer two") 
	defer fmt.Println("\nDefer three")
	//this statement will execute all the end of the func
	//in case of multiple defer statement in a func it will execute in LIFO way
	fmt.Println("Hello")
	 
	myDeffer()
}

func myDeffer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// Output:

// Hello
// 43210
// Three 
// Two
// One