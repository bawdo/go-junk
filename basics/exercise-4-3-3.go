// Go Bootcamp: 4.3.3
package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	var maxLen int
	for _, name := range names {
		if l := len(name); l > maxLen {
			maxLen = l
		}
	}
	output := make([][]string, maxLen)
	for _, name := range names {
		output[len(name)-1] = append(output[len(name)-1], name)
	}
	fmt.Printf("%v", output)
}
