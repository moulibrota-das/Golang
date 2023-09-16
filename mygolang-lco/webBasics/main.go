package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	res, err := http.Get(url)
	// using the Get() method to get the response from the specified url

	if(err != nil){
		panic(err);
	}

	fmt.Println(res);

	defer res.Body.Close() 
	//caller's responsibility to close the response

	data, err := io.ReadAll(res.Body)
	// using ioutils to read the http response and store it as bytedata
	if(err != nil){
		panic(err)
	}

	content := string(data)
	// converting the byte data to string for better understanding
	fmt.Println("Converted content is: \n",content)
}