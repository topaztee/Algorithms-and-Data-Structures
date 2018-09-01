package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := []int{1,2,3,4,5}
	rand.Seed(time.Now().UnixNano())
	for i := range deck {
		randCardIndex := rand.Intn(len(deck))
		fmt.Println(randCardIndex)
		deck[i], deck[randCardIndex] = deck[randCardIndex], deck[i]
		fmt.Println(deck)
	}
}
