package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza: ")

	//comma ok or err err syntax

	input, _ := reader.ReadString('\n') //when we dont care about errors, we just use underscore
	//input, err := reader.ReadString('\n') //when we care about errors and want them to be returned
	fmt.Println("Thanks for rating", input)
	fmt.Printf("type of this rating is  %T", input) //returns string
	//now as it returns string, we cn only perform string operations on it, so niw we need to know about conversions
}