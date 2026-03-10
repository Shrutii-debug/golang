package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in golang")

	rand.Seed(time.Now().UnixNano()) //we have seeded so that we can get random numbers every single time

	diceNumber := rand.Intn(6) + 1 //+1 because range is always not inclusive so 6 will be not inclusive therefore +1 to make it inclusive
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll the dice again")

	default:
		fmt.Println("what was that ?! ")
	}

}