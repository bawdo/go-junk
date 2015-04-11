// Go Bootcamp: 3.4

package main

import (
  "fmt"
   "time"
)

type Bootcamp struct {
  // Latitude of the event
  Lat float64
  // Longitude of the event
  Lon float64
  // Date of the event
  Date time.Time
}

func main() {
  fmt.Println(Bootcamp{
                Lat: 34.012836,
                Lon: -118.495338,
                Date: time.Now(),
  })
}
