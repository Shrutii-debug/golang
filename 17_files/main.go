package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//writing in a file
func main() {
	fmt.Println("welcome to files")

	content := "this needs to go inn a file - learncodeonline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil{
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil{
		panic(err) //panic will stop the execution process and tell about the error
	}

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}


//for file creation we use the os operation, but for reading and doing manipulation
//there is a separate utility called ioutil

//reading a file  

func readFile(filename string){
dataByte, err := ioutil.ReadFile(filename) //whenever a file is rread, it is not read in the format of data type provided
//a file is read in bytes

/*
if err != nil{
panic(err)
}
	*/

	checkNilErr(err)

	//fmt.Println("Text data inside the file is: \n", dataByte)
	//now this databyte will return a whole lotta diff things, so we can just wrap it around in a string

	fmt.Println("Text data inside the file is: \n", string(dataByte))
}

//now instead of using if else repeatedly, we can create a check func and use it

func checkNilErr(err error){
	if err != nil{
	panic(err)
	}
}
//now you can use this func