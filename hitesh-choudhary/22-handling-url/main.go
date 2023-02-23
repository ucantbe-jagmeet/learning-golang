package main

import (
	"fmt"
	"net/url"
)


const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=abcd456efg"


func main() {
	fmt.Println("Welcome to Handling URL in golang")
	fmt.Println(myurl)


	// parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)   // scheme is https
	fmt.Println(result.Host)	// host is lco.dev:3000
	fmt.Println(result.Path)	// path is /learn
	fmt.Println(result.Port())	 // port is 3000
	fmt.Println(result.RawQuery)	// rawquery is coursename.....

	qparams := result.Query()  // query stores Query parameters in a better format than RawQuery
	fmt.Printf("Thw type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])


	for _, val := range qparams{
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)



}