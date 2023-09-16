package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("Please rate our pizza app")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newInput := strings.TrimSpace(input)

	fmt.Println("Thanks for rating", newInput)

	numRating, err := strconv.ParseFloat(newInput, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating + 1)
	}
}