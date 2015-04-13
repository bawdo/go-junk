// Go Bootcamp: 4

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [2]string{"hello", "world!"}
	fmt.Println(b)
	fmt.Printf("s s%\n", b)
	// [hello world!]
	fmt.Printf("q q%\n", b)
	// ["hello" "world!"]

	var c [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("%q", c)
}
