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
	// EncodedJson()
	DecodeJson()
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

func DecodeJson(){
	 // imp-> any data that comes from the web is in the byte format , so we have to wrap it into string .
	 // sometimes we wrap string into bytes and sometimes we dont need this 

	 jsonDataFromWeb := []byte(`
	 {
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	 }
	 `)

	 var lcoCourse course
	 checkValid := json.Valid(jsonDataFromWeb)

	 if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)      //we have to pass lcocourse address because method 
															//unmarshal() can make copy of it 
		fmt.Printf("%#v\n",lcoCourse)
	 }else{
		fmt.Println("JSON WAS NOT VALID")
	 }

	 // some cases where you just want to add data to key value

	 var myOnlineData map[string]interface{}
	 json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	 fmt.Printf("%#v\n", myOnlineData)

	 for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type of the data is: %T\n ", k,v,v)
	 }

}