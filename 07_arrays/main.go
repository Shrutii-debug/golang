//as compared to other programming languages, arrays is least used in golang

package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	//in this case as we have initialized only 3 spaces but size is 4
	//so there wil be small space between apple and tomato but a much larger space between tomato and peach as comapred to apple and tomato
	fmt.Println("fruit list is :",fruitList)

	fmt.Println("fruit list is :",len(fruitList)) //this returns 4 even if 3 spaces are filled
	//because we have defined 4 in the memory, it does not count the content


	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("vegy list is:", vegList)
}