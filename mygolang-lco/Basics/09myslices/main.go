package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Peach"}

	fmt.Printf("The data type of fuitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Tomato")
	fmt.Println(fruitList)
	//inserting element in slices

	fruitList = append(fruitList[1:])
	fmt.Println("Slicing from 1 position to end", fruitList)
	//slicing from 1 pos to end

	highscores := make([]int, 4) //will make a slice of size 4
	highscores[0] = 121
	highscores[1] = 324
	highscores[2] = 651
	highscores[3] = 427
	// highscores[4] = 720 	this will give error arrayIndexOutOfBound

	highscores = append(highscores, 823, 849, 960)
	//this will not give error instead assign a new memory for the var


	sort.Ints(highscores)
	//sort integer slices

	fmt.Println(highscores)

	//How to remove element based on index

	var courses = []string{"react", "node", "next", "java"}
	var index int = 2
	fmt.Println(courses)

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}