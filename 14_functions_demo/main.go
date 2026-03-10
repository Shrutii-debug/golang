package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")
	greeter()

	result := adder(3, 5)

	fmt.Println("result is : ", result)

	//proResult := proAdder(2, 5, 8, 7)

	//like we have made the func to return int and string both, we also have to handle that
	//so we will declare a second variable or we can also write underscore if we dont care about it
	proResult, myMessage := proAdder(2, 5, 8, 7)


	fmt.Println("pro result is: ", proResult)
	fmt.Println("pro message is:", myMessage)

}
// you cant define a function inside  a function
//you have to define it outside and refer it 
func greeter(){
	fmt.Println("namastey from golang")
}

func adder(valOne int, valTwo int) int{ //this int outside the () shows the  data type of the value that youre going to return
	return valOne + valTwo
}

//i have no idea about how many values are coming, but I want to add all of them
//lets create a function for that

//syntax
/*
func proAdder(values ...int) int{
	total := 0

	for _, val := range values{
		total += val
	}
	return total
}
	*/

//if you want to return a string along with the int result you can do that too
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values{
		total += val
	}
	return total, "Hi pro result function"
}