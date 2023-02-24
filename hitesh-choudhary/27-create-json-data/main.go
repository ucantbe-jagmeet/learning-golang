package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int 	
	Platform string		`json:"website"`
	Password string		`json:"-"`
	Tags []string  		`json:"tags,omitempty"`  // omitempty is says if the value is nil then dont throw that field at all
}


func main() {
	fmt.Println("Welcome to creating Json data in golang")
	EncodedJson()
}

func EncodedJson(){
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc@123", []string{ "web-dev", "JS"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd@123", []string{ "full-stack", "JS"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "ang@123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent( lcoCourses, "", "\t")
	if err!=nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}