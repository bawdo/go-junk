// Go Bootcamp: 3.5

package main

import "fmt"

type Bootcamp struct {
  Lat, Lon float64
}

func main() {
  x := new(Bootcamp)
  y := &Bootcamp{}
  fmt.Println(*x == *y)
  fmt.Printf("%s == %s\n", x, y)
}
