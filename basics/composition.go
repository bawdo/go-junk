// Go Bootcamp: 3.6

package main

import "fmt"

type User struct {
  Id       int
  Name, Location string
}

func (u *User) Greetings() string {
  return fmt.Sprintf("Hi %s from %s",
           u.Name, u.Location)
}

type Player struct {
  User
  GameId int
}

func main() {
  // dot notation to set fields
  p := Player{}
  p.Id = 42
  p.Name  = "Bawdo"
  p.Location  = "VIC"
  p.GameId = 100001
  fmt.Printf("%+v", p)
  fmt.Println(p.Greetings())

  // struct literal use to initialize
  p2 := Player{
          User{Id: 43, Name: "Bawdo2", Location: "VIC"},
          10001,
        }
  fmt.Printf("Id: %d, Name: %s, Location: %s, Game id: %d\n",
        p2.Id, p2.Name, p2.Location, p2.GameId)
  // Directly define a field on the Player struct
  p2.Id = 11
  fmt.Printf("%+v\n", p2)
  fmt.Println(p2.Greetings())
}
