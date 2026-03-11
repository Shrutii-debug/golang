package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("get request")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

//GET request in golang

func PerformGetRequest(){
	const myurl = "https://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())


	//fmt.Println(content)
	//fmt.Println(string(content))
}

//Making a post request
func PerformPostJsonRequest(){
	const myurl = "https://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang",
			"price" : 0,
			"platform" :"learncodeonline.in"

		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody) //all the three parameters

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
 

//Post form operation

func PerformPostFormRequest(){
	const myurl = "https://localhost:8000/postform"

	//form data

	data := url.Values{}
	data.Add("firstname", "shruti")
	data.Add("lastname", "agarwal")
	data.Add("email", "shruti@go.dev")

	response, err := http.PostForm(myurl, data) //in this we dont have to mention the format in which our data is coming unlike JSON

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}