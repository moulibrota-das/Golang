package main

import "fmt"

//struct name starting with Caps letter means it is public outside this package
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {

	fmt.Println("Structs in Golang")

	user1 := User{Name: "Shinchan",Email: "balleballeshabashaba@nohara.fam", Age: 5, Status: false}
	fmt.Println(user1)
	fmt.Printf("%+v\n", user1)
}