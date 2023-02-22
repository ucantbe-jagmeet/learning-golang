// Simple dice game

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Dice game using switch case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) +1

	fmt.Println("value on dice is ", diceNumber)

	switch diceNumber {
	case 1: fmt.Println("Dice value is 1 and you can open")
	case 2: fmt.Println("You can move 2 Spot")
	case 3: fmt.Println("You can move 3 Spot")
			fallthrough
	case 4: fmt.Println("You can move 4 Spot")
			fallthrough
	case 5: fmt.Println("You can move 5 Spot")
	case 6: fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!!!!")
	}
}