package main

import "fmt"

const LogInToken string = "hI there !!"; // public variable always has capital letter

func main() {
	var username string = "Jagmeet Singh";
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n", username);
	
	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type: %T \n", isLoggedIn);
	
	
	var smallVal uint8 = 255;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type: %T \n", smallVal);
	
	var smallFloat float64 = 255.424242131231290118519;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n", smallFloat);
	
	
	//default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type: %T \n", anotherVariable);

	//implicit type
	var website = "thisistrending.in";
	fmt.Println(website)

	// no var style
	numberOfUser := 30000;   // ( := volurous operator)
	fmt.Println(numberOfUser);

	fmt.Println(LogInToken);
	fmt.Printf("Variable is of type: %T \n", LogInToken);

}