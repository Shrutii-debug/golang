// structs are very important as they are the alternative for clauses, we dont have clauses in golang

package main

import "fmt"

func main()  {
	fmt.Println("welcome to structs")
	// there is no inheritance in golang, there is no concept of super or parent

	
	shruti := User{"Shruti", "shruti@go.dev", true, 5}
	fmt.Println(shruti)

	// using the +v syntax, it gives the whole structure, entire details along with the parameters
	fmt.Printf("shruti details are: %+v\n", shruti)

		//if you want to get specific details

	fmt.Printf("name is %v and email is %v.", shruti.Name, shruti.Email) //us ecapital letters as they are being exported

}

//defining a struct

type User struct {
	Name string
	Email string
	Status bool
	Age int
}