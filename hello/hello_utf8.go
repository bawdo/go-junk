package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("名前？: ")
  text, _ := reader.ReadString('\n')
  text = strings.TrimRight(text, "\n")
  fmt.Printf("こにちは%sさん", text)
}
