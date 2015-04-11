// Go Bootcamp: 3.4

package main

import "fmt"

type Point struct {
  X, Y int
}

var (
  p = Point{1, 2}  // has type Point
  q = &Point{1, 2} // has type *Point
  r = Point{X: 1}  // Y:0 is implicit
  s = Point{}      // X:0 and Y:0
)

func main() {
  fmt.Println(p, q, r, s)
}
