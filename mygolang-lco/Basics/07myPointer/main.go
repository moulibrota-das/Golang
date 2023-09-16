package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var one int = 2
	var ptr *int = &one// pointer syntax;
	var ptr1 = &one

	fmt.Println("value of ptr is ", ptr);
	fmt.Println("value of ptr1 is ", ptr1);
	fmt.Println("ptr is storing value ", *ptr);

}