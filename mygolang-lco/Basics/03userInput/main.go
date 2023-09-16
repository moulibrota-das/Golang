package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {
	welcome := "Welcome to user input"

	fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our lemon tea")

	//comma ok || err ok

	// input, _ := reader.ReadString()
	var input int
	fmt.Scan(&input)
	fmt.Printf("Thanks for rating %d stars\n", input)
	fmt.Println(input)
}