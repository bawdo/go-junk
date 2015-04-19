// Go Bootcamp: Ch 7

package main

import (
	"fmt"
	"os"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

// This interface is satisfied by types that implement the Name() method.
type Namer interface {
	Name() string
}

// Namer must have the Name() defined.
func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

// Interfaces can contain other interfaces
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	u := &User{"Keith", "Bawden"}
	fmt.Println(Greet(u))

	c := &Customer{42, "Frencesc"}
	fmt.Println(Greet(c))

	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "Hello, writer\n")
}
