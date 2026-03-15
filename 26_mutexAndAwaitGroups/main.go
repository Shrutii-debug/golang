package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("race condition")

	wg := &sync.WaitGroup{}
	//mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	var score = []int{0} //0 is the default value

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("one R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("three R")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}

//use go run --race to know about the race condition