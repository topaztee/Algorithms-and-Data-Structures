package main

import (
	"fmt"
	"strings"
	"time"
)
// Two goroutines running at the same time
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%sf(%d): %d\n", strings.Repeat(" ", n*4), n, i)
		time.Sleep(1 * time.Microsecond)
	}
}

func spawnTen() {
	for i := 0; i < 10; i++ {
		go func(n int){
			fmt.Println("Goroutine", n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	
}

func main() {
	// go f(0)
	// f(1)

	//spawn 10 goroutines
	spawnTen()

	// spawnInALoopFixed()
}
