// Go Bootcamp: 3.3

package main

import (
  "fmt"
  "time"
)

func timeMap(y interface{}) {
  z, ok := y.(map[string]interface{})
  if ok {
    z["updated_at"] = time.Now()
  }
}

func main() {
  foo := map[string]interface{}{ "Bawdo": 41, }
  timeMap(foo)
  fmt.Println(foo)
}
