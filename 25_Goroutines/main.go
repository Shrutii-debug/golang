package main

import (
	"fmt"
	"net/http"
	"sync"
	//"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup //usually these are pointers

var mut sync.Mutex //pointer

/*
1️⃣ The Problem WaitGroups Solve

In Go, when you start goroutines, they run asynchronously.

Example:

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello")
}

func main() {
	go sayHello()
	fmt.Println("Main finished")
}
Output
Main finished

You might not see "Hello".

Why?

Because:

main() finishes
↓
program exits
↓
goroutine is killed

So the goroutine didn't get time to finish.

2️⃣ What WaitGroup Does

A WaitGroup makes the main program wait for goroutines to finish.

Conceptually:

Start tasks
↓
Wait until all tasks finish
↓
Continue program
3️⃣ WaitGroup Comes From

Package:

import "sync"

Structure:

var wg sync.WaitGroup
4️⃣ The Three Important WaitGroup Functions
Function	Meaning
Add(n)	number of goroutines to wait for
Done()	marks one goroutine finished
Wait()	blocks until all goroutines finish
5️⃣ Basic Example
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("Worker", id, "starting")

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	wg.Wait()

	fmt.Println("All workers finished")
}
6️⃣ Execution Flow
main starts
↓
wg.Add(3)
↓
3 goroutines start
↓
each calls wg.Done()
↓
counter goes 3 → 2 → 1 → 0
↓
wg.Wait() unblocks
↓
program continues
7️⃣ Visual Intuition

Think of WaitGroup like a task counter.

wg.Add(3)

Counter = 3

When each goroutine finishes:

wg.Done()

Counter = Counter - 1

When counter becomes 0:

wg.Wait() stops waiting
8️⃣ Very Important: Pass WaitGroup by Pointer

Correct:

func worker(wg *sync.WaitGroup)

Incorrect:

func worker(wg sync.WaitGroup)
Why?

Because if you pass by value:

goroutine gets a copy

The original counter won’t change.

Pointer ensures:

all goroutines modify same WaitGroup
9️⃣ Real Backend Example

Imagine calling multiple APIs simultaneously.

API 1
API 2
API 3

Instead of waiting sequentially:

call API1 → wait
call API2 → wait
call API3 → wait

You run them in parallel:

wg.Add(3)

go callAPI1()
go callAPI2()
go callAPI3()

wg.Wait()

This makes backend much faster.

🔟 Important Mistake Beginners Make

Wrong:

wg.Add(1)

go func() {
	wg.Add(1)   // ❌ wrong place
}()

Add() should be called before starting goroutines.

Correct:

wg.Add(1)
go worker()
11️⃣ Another Best Practice

Often you will see:

defer wg.Done()

Example:

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("working")
}

Why?

Even if an error happens, Done() will run.

12️⃣ WaitGroup vs Channels (Important)

Go concurrency mainly uses two tools:

Tool	Purpose
WaitGroup	wait for goroutines
Channels	communicate between goroutines
13️⃣ Real System Example

Suppose a backend request needs:

User info
Orders
Notifications

Instead of fetching sequentially:

User → Orders → Notifications

You do:

go getUser()
go getOrders()
go getNotifications()

wg.Wait()

This makes the API 3× faster.

Key Takeaway

WaitGroup ensures:

Program waits for all goroutines to finish before continuing.

Without it:

goroutines may terminate early.
*/

func main() {
	//goroutine is created by adding a keyword go
	//go greeter("hello") //after adding go, only world is printed 5 times
	//greeter("world")
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}

	for _, web := range websitelist{
		getStatusCode(web)
		wg.Add(1) //keeps on adding the go routines number ofhow many of my frnds are out here
	}
	wg.Wait() //always comes in the end of main func which species it to wait and not 
	fmt.Println(signals)
}

/*
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
	*/

func getStatusCode(endpoint string){
	defer wg.Done()


	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in ednpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf(" %d status code for %s\n", res.StatusCode, endpoint)
	}
}