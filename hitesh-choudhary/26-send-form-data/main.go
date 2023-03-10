package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)


	var responseString strings.Builder

	content, _ :=  ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println( "Byte count is: " ,byteCount) // shows content length
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(` 
	{
		"courseName": "Lets go with golang",
		"price": 0,
		"platform": "learnCodeonline.in"
	}

	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "Jagmeet")
	data.Add("lastname", "Singh")
	data.Add("email", "jagmeet@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}