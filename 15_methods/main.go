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

	fmt.Printf("name is %v and email is %v. \n", shruti.Name, shruti.Email) //use ecapital letters as they are being exported

	shruti.GetStatus()
	shruti.NewMail()

	fmt.Printf("name is %v and email is %v. \n", shruti.Name, shruti.Email)
	//original email remains the same even if we have mentioned a diff email in the function
	
}

//defining a struct

type User struct {
	Name string
	Email string
	Status bool
	Age int
	
}

//defining a method
func (u User) GetStatus() {
	fmt.Println("is user active?", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("email of this user is: ", u.Email)
}