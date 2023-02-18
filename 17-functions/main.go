// we can add functions in another function

package main

import "fmt"

func main(){

	fmt.Println("Welcome to functions")
	greet()

	result := adder(3,5)
	fmt.Println("REsult is,",result)
	
	proResult, myMessage := proAdder(2,4,5,6,7,8,9,10)
	fmt.Println("Pro result is,",proResult)
	fmt.Println("My message is,",myMessage)


}

func adder( valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values{
		total += val
	}

	return total, "Hi your result is good"
}

func greet(){
	fmt.Println("Namastey from Golang")
}

func greetTwo () {
	fmt.Println("Another method greet")
}

