/*

 */

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in go lang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mygofile.txt")  // os is used to create the file

	checkNilErr(err)


	length, err := io.WriteString(file, content) // io is used to write anything
	checkNilErr(err)


	fmt.Println("Length is:",length)
	defer file.Close()

	readFile("./mygofile.txt")


}

func readFile(fileName string){
	databyte, err := ioutil.ReadFile(fileName)     // ioutil is used to read the file
	checkNilErr(err)

	fmt.Println("text data inside the file is \n",string(databyte))

}


func checkNilErr(err error){
	if err!=nil{
		panic(err) // panic will shut down the program and throw whatever error we get
	}
}