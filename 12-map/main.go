/*
	new() -> it swill give alot of errors, we cant put any values in it because it the zeroed started value
	so we use make()

*/

package main

import "fmt"

func main(){
	fmt.Println("Welcome to video on map function")

	langauges := make(map[string]string)

	langauges["JS"] = "Javascript"
	langauges["RB"] = "Ruby"
	langauges["PY"] = "Python"

	fmt.Println("List of ALl langauges: ",langauges)
	fmt.Println("JS shorts for: ",langauges["JS"])
	fmt.Println("RB shorts for: ",langauges["RB"])
	fmt.Println("PY shorts for: ",langauges["PY"])
	
	delete(langauges, "RB")
	fmt.Println("List of ALl langauges: ",langauges)
	
	//loops are interesting  in go long

	for key , value := range langauges{
		fmt.Printf("For Key %v, value is %v\n", key, value )
	}		

}