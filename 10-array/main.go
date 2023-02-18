package main

import "fmt"

func main(){

	fmt.Println("Welcome to array in go lang")

	var list [4]string

	list[0] = "Apple"
	list[1] = "Orange"
	// list[2] = "Grapes"
	list[3] = "Peach"

	fmt.Println("Fruit list is ", list)
	fmt.Println("Fruit list is ", len(list))

	var list1 = [3]string{ "potato", "beans", "mushroom"}
	fmt.Println("Veggie list is", list1)
	fmt.Println("Veggie list is", len(list1))
}