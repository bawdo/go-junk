// Go Bootcamp: 5.4

package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	scores       = map[string]int{
		"a": 1,
		"e": 1,
		"i": 2,
		"o": 3,
		"u": 4,
	}
)

func CharCoins(char string) int {
	return scores[strings.ToLower(char)]
}

func NameCoins(name string) int {
	score := 0
	for _, char := range name {
		score += CharCoins(string(char))
	}
	if score > 10 {
		return 10
	} else {
		return score
	}
}

func main() {
	for _, user := range users {
		score := NameCoins(user)
		distribution[user] += score
		coins -= score
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
