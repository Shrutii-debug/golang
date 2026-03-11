package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=adwi648jns"



func main() {
	fmt.Println("welcome to urls")
	fmt.Println(myUrl)


	//parsing the url, meaning extracting some information

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme) //returns https
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) //rest all are properties, but this port is a method not a property
	fmt.Println(result.RawQuery)

	//another way of extracting information and this is a better way
	//qparams means query parameters

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"]) //returns reactjs

	for _, val := range qparams{
		fmt.Println("param is: ", val)
	}


	//okay the second part, maybe you have things in chunks and you want to construct a url out of it

	//this is very imp, you always pass a reference of url, not a copy of it
	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=shruti",
	}
	anotherURL := partsofUrl.String() //another way of using and converting something to a string
	// we could have done it manually too
	fmt.Println(anotherURL)
}