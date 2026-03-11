package main

import (
	
	"fmt"
	"io"
	"net/http"
)


const url = "https://lco.dev"


func main() {
	fmt.Println("welcome to web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type:  %T \n", response)

	defer response.Body.Close() //it is caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}