package main

import "fmt"

func main() {

	//compulsary mention the size of array
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is ", fruitList);

	var vegList = [3]string{"potato", "peas", "mushroom"}

	fmt.Println("Veg list is ", vegList)
}