// Go Bootcamp: 8.2.1

package main

import "fmt"

func main() {
	c := make(chan int, 2) // blocked until 2 ints have been recieved
	c <- 1
	c <- 2
	// c <- 3 Doing this would lead to the following error:
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c)
	fmt.Println(<-c)

	d := make(chan int, 2)
	d <- 1
	d <- 2
	d3 := func() { d <- 3 }
	go d3() // Does not block the main thread.
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
}
