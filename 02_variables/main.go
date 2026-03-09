package main


import "fmt"

const LoginToken string = "jbkBjher"
//putting up first letter as capital letter makes the variable a public variable
//so it is almost similar to adding a Public keyword out there
//this variable can be used anywhere in the program

func main() {
	var username string = "shruti"
	fmt.Println(username)
	//fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	//fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	//fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4556738911234
	fmt.Println(smallFloat)
	//fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) //when we do ont initialize, the default value is 0
	//fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type

	var website = "github.com"
	fmt.Println(website)

	//lexar comes and decides the type based on the content inside the variable
	// but if we dont declaare then we cant change the content later on such as doing website = 3

	//no var style
	numberOfUser := 30000 //inside any method we can use this
	//but we cant use outside in a global state
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)


}