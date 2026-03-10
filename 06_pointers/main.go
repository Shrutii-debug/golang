//A pointer is a variable that stores the memory address of 
// another variable. Instead of holding the actual value, it points to where that value is stored in memory.

package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	//just how we create normal variables we create pointers like that
	//just use * before datatype
	//now this means that the pointer will hold values of integer data type


	//var ptr *int //default value is <nil>

	//fmt.Println("value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber //creating a pointer which is referencing something
	//while referencing we use & to refer, and while just cfeating we use *

	fmt.Println("value of actual pointer is ", ptr) //this returns memory address
	fmt.Println("value of actual pointer is ", *ptr) //this returns the value inside the pointer which is 23

	//basically *ptr is the value inside the memory address
	*ptr = *ptr * 2
	fmt.Println("new value is", myNumber) //returns 46
	// so pointers guarantees that all the operations take place on actual values and not on the copies
	//original values are changed
}