/*
	when ever you pass a variable, it passess a copy of that variable
	Pointers is nothing its just memory reference of that variable
	* is used to creating a pointer
	& is used to referencing a pointer
*/

package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer is ", ptr)

	myNum := 23

	var ptr  = &myNum

	fmt.Println("Value of acutal pointer is: ", ptr) // it shows the reference of memory location
	fmt.Println("Value of acutal pointer is: ", *ptr) //  it show the value assigned at that refernce point

	*ptr = *ptr * 2
	fmt.Println("New value is : ", myNum)
}