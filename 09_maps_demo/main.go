package main

import (
	"fmt"

	
)

func main() {
	fmt.Println("maps in golang")

	//just like we can use make() to create slices, we can also use it to create maps
	//but we have to define that we need a map and the key and value data type

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)

	fmt.Println("JS shorts for: ", languages["JS"])

	//deleting values, you can use this for slices as well

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	//looping through maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value) //%v is used for values, %T for types, see more in documentation
	}

	// as we have the := operator, so incase you dont care about the keys
	//you can just remove it and put underscore
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n",  value) //%v is used for values, %T for types, see more in documentation
	}

}