package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	mapCount := make(map[string]int)
	wordsArr := strings.Fields(s)
	for _, i := range wordsArr {
		mapCount[i]++
	}
	return mapCount
}

func main() {
	// wc.Test(WordCount)
	fmt.Println(WordCount("hello my name is hello"))
}
