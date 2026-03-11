package main

import "fmt"

/*
defer Keyword in Go

In Go, the defer keyword is used to delay the execution of a function until the surrounding function finishes.

It is commonly used for cleanup tasks like:

closing files

releasing resources

unlocking mutexes

logging

Think of it as:

"Run this statement just before the function returns."

1️⃣ Basic Example
package main
import "fmt"

func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
Output
Hello
World

Explanation:

defer fmt.Println("World") is registered.

fmt.Println("Hello") runs first.

When main() finishes, deferred statement runs.

2️⃣ Multiple Defer Statements (LIFO Order)

Deferred calls execute in Last In First Out (LIFO) order.

package main
import "fmt"

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")

    fmt.Println("Start")
}
Output
Start
3
2
1

The last defer runs first.

3️⃣ Real Use Case: Closing Files

Very common in backend development.

package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error opening file")
        return
    }

    defer file.Close()

    fmt.Println("File opened successfully")
}

Why use defer here?

Because no matter how the function exits, the file will be closed.

4️⃣ Defer with Functions
package main
import "fmt"

func main() {
    defer greet()
    fmt.Println("Main function running")
}

func greet() {
    fmt.Println("Hello from defer")
}

Output:

Main function running
Hello from defer
5️⃣ Defer Arguments Are Evaluated Immediately

Important concept.

package main
import "fmt"

func main() {
    x := 10
    defer fmt.Println(x)

    x = 20
}

Output:

10

Why?

Because x was evaluated when defer was declared, not when it executed.

6️⃣ Defer with Anonymous Functions
package main
import "fmt"

func main() {
    x := 10

    defer func() {
        fmt.Println(x)
    }()

    x = 20
}

Output:

20

Because the anonymous function reads the current value when executed.

7️⃣ Real Backend Example

Example in a web server where database connection must close.

db, err := sql.Open("mysql", connectionString)
if err != nil {
    log.Fatal(err)
}

defer db.Close()

This guarantees the DB connection closes after function execution.

⚡ Key Points
Feature	Explanation
Execution time	Runs when function returns
Order	LIFO (stack)
Common uses	closing files, DB connections, locks
Arguments	evaluated immediately

💡 Interview Tip (important)
A common Go interview question: What happens if defer is inside a loop?
Example:  for i := 0; i < 3; i++ {
    defer fmt.Println(i)
}
Output
2
1
0

Because defer calls stack up and execute after the function ends.
*/

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	myDefer()

	//defer fmt.Println("world")
}

func myDefer(){
	for i := 0; i < 5; i++{
		defer fmt.Println(i)
	}
}