package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 { //this curly braces cant be moved to the next line
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10 login counts"
	}


	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	//you can initialize a variable and directly put a value into it
	if num := 3; num < 10{
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	/*
	  if err != nil {

	}
*/
}