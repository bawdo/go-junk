// Go Bootcamp: 2.7

package main

import "fmt"

type Artist struct {
  Name, Genre string
  Songs int
}

func newRelease(a Artist) int {
  a.Songs++
  return a.Songs
}

// Takes a pointer value
func newRelease2(a *Artist) int {
  a.Songs++
  return a.Songs
}

func main() {
  me := Artist{Name: "Bawdo", Genre: "Geek Folk", Songs: 43}
  fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
  fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

  // Get the pointer value for Artist.
  me2 := &Artist{Name: "Bawdo", Genre: "Geek Folk", Songs: 43}
  fmt.Printf("%s released their %dth song\n", me2.Name, newRelease2(me2))
  fmt.Printf("%s has a total of %d songs\n", me2.Name, me2.Songs)
}
