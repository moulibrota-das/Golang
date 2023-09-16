package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// func proAdd(values ...int) int{
// 	total := 0

// 	for _, value := range values{
// 		total += value
// 	}

// 	return total
// }

func proAdd(values ...int) (int, string){
	total := 0

	for _, value := range values{
		total += value
	}

	return total,"String return message"
}


func main() {
	result := add(38 , 2)
	proResult,msg := proAdd(1,2,3,4,5,6,7,8)

	fmt.Println(result)
	fmt.Println(proResult, msg)
}