// Go Bootcamp: 6.2

package main

import (
	"fmt"
	"math"
	"strings"
)

// We cannot extend the basic string type directly so we need to create an
// alias of it and extend that instead.
type MyStr string

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	fmt.Println(MyStr("test").Uppercase())
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs)
}
