// comma ok syntax

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input";
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating for our Pizza: ");

	// comma ok || err ok 

	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for Rating, ", input);

	fmt.Printf("Type of rating is Rating %T ", input);
}