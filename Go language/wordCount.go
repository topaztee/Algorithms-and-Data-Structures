package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	mapCount := make(map[string]int)
	wordsArr := strings.Fields(s)
	for v, i := range wordsArr {
		mapCount[v]++
	}
	return mapCount
}

func main() {
	// wc.Test(WordCount)
	fmt.Println(WordCount("hello my name is hello"))
}
