package main

import "fmt"

func main() {
	fmt.Println("Welcome to lecture on Structs")
	// no inheritance in golang
	// no super or parent concepts

	jagmeet := User{ "Jagmeet Singh", "jagmeets@gmail.com", true, 23}

	fmt.Println(jagmeet)

	fmt.Printf("Jagmeet details are: %+v\n",jagmeet)
	fmt.Printf("Name is %v and email is %v",jagmeet.Name, jagmeet.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age	int
}