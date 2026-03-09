package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")
//using the time package
	prsentTime := time.Now()
	fmt.Println(prsentTime)


//this is the predefined syntax, we always have to use these values
	fmt.Println(prsentTime.Format("01-02-2006 15:04:05 Monday"))
	//if you want just the date or the time just type that and remove all others

	//month is of type time.Month rest all are integers , time is in nanosecond and so on
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
	//this looks bad so weneed to format it

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

//now if you want to build something that shows the time on command line 
//go allows you to build
// GOOS="windows" go build //this is the command //this will create a file of mytime.exe