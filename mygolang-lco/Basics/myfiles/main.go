package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	content := "This a example content for the file created using golang"

	file, err := os.Create("./exampleFile.txt")
	//create a file in current dir naming exampleFile.txt

	if(err != nil){
		panic(err)
	} 
	
	length, err := io.WriteString(file, content)
	// write string in a existing file

	// if(err != nil){
	// 	panic(err)
	// }
	checkNilError(err);

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./exampleFile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err);

	fmt.Println("byte file is: ",databyte)
	fmt.Println("string file is: ",string(databyte))
}

func checkNilError(err error){
	if(err != nil){
		panic(err)
	}
}