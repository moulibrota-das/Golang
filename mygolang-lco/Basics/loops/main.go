package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")

	// arr := []string{"Gandu", "Bokachoda", "Chutiya", "Gere", "Khanki"}

	// for i:=0; i<len(arr); i++{
	// 	fmt.Println(arr[i]);
	// }

	// for i:= range arr{
	// 	fmt.Println(arr[i])
	// }

	// for index, element := range arr{
	// 	fmt.Println(element)
	// }

	// for _, element := range arr{
	// 	fmt.Println(element)
	// }

	n := 1

	for n <= 10{
		fmt.Println(n);
		n++;
	}
}