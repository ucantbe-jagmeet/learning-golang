/*
A 'defer' statement invokes a function whose execution is deferred to the moment the surrounding
funciton returns

when a function executes it executes line by line that what the flow of the function
but when we put differ keyword whatever the next line is written in function , it will execute
at the very end of the function itself

deferred functions are invoked immediately before the surrounding function returns, in the
reverse order they were deferred ( LIFO )

*/

package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++{
		defer fmt.Println(i)
	}
}