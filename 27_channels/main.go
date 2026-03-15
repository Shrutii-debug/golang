//channels are a way in which multiple goroutines can talk to each other
//they will still not be aware of what is happening inside that
//or how long its going to take another goroutine to finish its task

package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("channels in golang")

	myCh := make(chan int, 2)
	//we can use this 2 which means we are usign buffered chanel
	wg := &sync.WaitGroup{}

	//myCh <- 5
	//fmt.Println(<-myCh)
	wg.Add(2)
	//<-chan --> receive only, this thing then doesnt allow to use Close
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	//<- send only
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh <- 0
		//myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}

