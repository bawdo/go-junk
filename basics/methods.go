// Go Bootcamp: 6
// Methods: "You can define a method on _any_ type in your your package, not
// just structs. You cannot define a method on a type from another package, or
// on a basic type". SOURCE: Go Bootcamp

package main

import (
	"fmt"
	"math"
)

type User struct {
	FirstName, LastName string
}

// User is the reciever of this method
// NOTE: the method is defined outside of the struct definition
func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

// Same method associated with a pointer to the User type instead.
// This avoids copying the User strunct on each call.
func (u *User) Greeting2() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

// Associating with a pointer also allows you to modify the value on the
// reciever that you are pointing to.
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	u := User{"Keith", "Bawden"}
	fmt.Println(u.Greeting())

	u2 := &User{"Keith", "Bawden"}
	fmt.Println(u2.Greeting())

	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
