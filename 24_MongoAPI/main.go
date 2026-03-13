package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shrutii-debug/mongoapi/routes"
)

func main() {
	fmt.Println("mongoDB api")
	r := routes.Router()
	fmt.Println("Server is getting started ...")
	//http.ListenAndServe(":4000", r)

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000...")

}