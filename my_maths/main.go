package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
	//"time"
)

func main() {
	fmt.Println("welome to maths")

	//var  myNumberOne int = 2
	//var  myNumberTwo float64 = 4.5

	//random number

	
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)

	//random number from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
	
}