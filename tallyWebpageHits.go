package main

import (
	"fmt"
)

func main() {
	type Key struct {
		path, country string
	}

	hits := make(map[Key]int)

	hits[Key{"/blog", "au"}]++
}