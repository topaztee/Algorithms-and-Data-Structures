package main

import (
	"regexp"
	"fmt"
)

func main() {
	spaces := lengthOfLastWord("Hello, my name is John")
	fmt.Println(spaces)
}

func lengthOfLastWord(s string) int {
	re := regexp.MustCompile(`\s`)
	t := re.FindAllString(s, -1)
	return len(t)
    
}

