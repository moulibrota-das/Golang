package main

import "fmt"

// jwtToken := 34913840 this is not allowed
// var jwtToken = 39283058 //this is allowed

const LoginToken string = "logintoken"	
// the first char of the variable name is in Caps which declare the variable as public


func main() {
	// var username string = "mouilbrota"
	// fmt.Println(username)
	// fmt.Printf("variable if of type: %T\n", username)

	// var isLoggedIn bool = false
	// fmt.Println(isLoggedIn)
	// fmt.Printf("variable if of type: %T\n", isLoggedIn)

	// var smallVal int = 255
	// fmt.Println(smallVal)
	// fmt.Printf("variable if of type: %T\n", smallVal)

	// var smallFloat float32 = 255.3134356
	// fmt.Println(smallFloat)
	// fmt.Printf("variable if of type: %T\n", smallFloat)

	//default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable if of type: %T\n", anotherVariable)

	//implicit type

	var website = "google.com"
	fmt.Println(website)

	//no var style

	numberOfUser := 6999
	fmt.Println(numberOfUser)
}