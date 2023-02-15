/*

we are not initialise the size of an array here
thus it make it more usefull to create dynamic array with use slices

*/

package main

import (
	"fmt"
	"sort"
)

func main(){

	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Peach", "Tomato"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3] ) // 1 and colon 3 which means staring and ending point and is used to make slice of the current slices
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 456
	highScore[3] = 867

	highScore = append(highScore, 555, 666, 777)

	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	
	sort.Ints(highScore) // these methods are present in slices and not in arrays
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))


	// how to remove value form slices based on index

	var courses = []string{ "ReactJs", "Javascript", "Ruby","Python", "Swift"}

	fmt.Println(courses)
	
	var index int =2;
	courses = append(courses[:index] , courses[index+1:]...)
	fmt.Println(courses)
} 