// Go Bootcamp: 5

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var answer int
	answer = 43

	if answer != 42 {
		fmt.Println("Wrong answer")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 5.3 Switch
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}

	num := 3
	v := num % 2
	switch v {
	case 0:
		fmt.Println("even")
	case 3 - 2:
		fmt.Println("odd")
	}

	rand.Seed(time.Now().Unix())
	score := rand.Intn(15 - 1)
	fmt.Println(score)
	switch score {
	case 0, 1, 2, 3:
		fmt.Printf("Your score %d is terrible\n", score)
	case 4, 5:
		fmt.Printf("Your score %d is mediocre\n", score)
	case 6, 7:
		fmt.Printf("Your score %d is not bad\n", score)
	case 8, 9:
		fmt.Printf("Your score %d is almost perfect\n", score)
	case 10:
		fmt.Printf("Your score %d is perfect. Did you cheat?\n", score)
	default:
		fmt.Printf("Your score %d is off the chart!\n", score)
	}

	// Switch with fallthroughs
	switch score {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	default:
		fmt.Println("Try again!")
	}

	// Switch with a break
	n := 1
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		if time.Now().Unix()%2 == 0 {
			fmt.Println("un pasito pa lante maria")
			break
		}
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
		fallthrough
	default:
		fmt.Println("Try again!")
	}

}
