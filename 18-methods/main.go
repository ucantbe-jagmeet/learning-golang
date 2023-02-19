// functions in the structs/classes refers to as method
// golang doesn't has classes in it so here we methods are initialised in structs

package main

import "fmt"

func main(){
	fmt.Printf("Welcome to methods in go lang\n\n")

	jagmeet := User{ "Jagmeet Singh", "jagmeet@go.dev", true, 16}
	fmt.Println(jagmeet)

	jagmeet.GetStatus()
	jagmeet.NewMail()
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
	//oneAge int // oneAge is not exportable as its first letter is not capital

}

func ( u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
	
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:",u.Email)
}

