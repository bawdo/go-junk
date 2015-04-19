// Go Bootcamp: Ch 7.2

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Anything that implements the Error() method will be implementing the
// builtin error interface.
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
