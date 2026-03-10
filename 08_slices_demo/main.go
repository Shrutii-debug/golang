package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	//var fruitList = []string{} //in this syntax of slices, you need to initialize it {}
	var fruitList = []string{"Apple", "Tomato", "Peach"} 
	//fmt.Printf("type of fruitlist is %T \n", fruitList) //returns []string

	//syntax for adding data or element
	fruitList = append(fruitList, "Mango", "Banana" )
	//fmt.Println(fruitList)


	//: is used to slice(Separate) the fruitlist
	//fruitList = append(fruitList[1:]) //i will start from 0, i will miss that 0, and ending limit is end
	fruitList = append(fruitList[1:3]) //range are always non inclusive so 0 will be missed, 1 and 2 will be printed, and3 will be missed
	fruitList = append(fruitList[:3]) //same 3 will be excluded, and it will start from 0 as nothing is mentioned
	//fmt.Println(fruitList)


	//using a diff syntax
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 //this will give error ofcourse as it is index out of bounds
	//but what happens if you use append method


	highScores = append(highScores, 555, 666, 321)//this will add the values to highscores
	//now append reallocates the memory and thus it is able to do this, whatever we have declared and initialized 
	//above, is the default allocation of memory
	//this is a function which is not available in array

	//now lets look at some other methods that is not available in the array



	sort.Ints(highScores) //sorts the slice in ascending order


	//fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //returns true if the values are sorted, otherwise false

	

	//how to remove a value from slices based on the index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2

	//you can use append to remove the values as well using the concept of colon 
	 //as we know that the last value will not be included, so index will not be included

	//courses = append(courses[:index], courses[index+1:])

	 //now this will give us error as we are passing more arguments
	// so we need to define func in a way such that it accepts more arguments
	// we will use ... for this

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // now swift is removed 

}