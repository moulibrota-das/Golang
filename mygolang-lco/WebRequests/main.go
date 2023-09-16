package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to handling Web req")
	// getMethod();
	// postMethod();
	performFormRequest();
}

func getMethod(){
	const url string = "http://localhost:8000/get";

	res, err := http.Get(url)
	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status Code is: ", res.StatusCode)
	fmt.Println("Content Length is: ", res.ContentLength)

	data,_ := io.ReadAll(res.Body)

	// fmt.Println("Content is: ", string(data))

	var responseString strings.Builder
	responseString.Write(data)
	// this return a byteCount and the store the data within responseString

	fmt.Println(responseString.String())
}

func postMethod(){
	const myurl string= "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"go with golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if (err != nil) {
		panic(err)
	}

	defer response.Body.Close()

	data,_ := io.ReadAll(response.Body)
	fmt.Println(string(data))
}

func performFormRequest(){
	const myurl string = "http://localhost:8000/postform";

	//form data

	data := url.Values{};

	data.Add("firstname", "Moulibrota")
	data.Add("lastname", "Das")
	data.Add("email", "moulibrota@gmail.com")


	//in this method we don't have to mention the content type.
	res, err := http.PostForm(myurl, data)

	if(err != nil){
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}