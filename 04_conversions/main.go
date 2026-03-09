package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 and 5: ")

	//taking user input

	reader := bufio.NewReader(os.Stdin)

	input, _ :=reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	//numRating = input + 1 //this gives error as input is of string type and it cant add number into a string
	//now we will convert this with the help of a package strconv

	//numRating, err := strconv.ParseFloat(input, 64)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}

	// even this returns an error saying that not only 4 but 4\n is also coming in the conversion
	//so we need to remove that, we will use strings package for it

}