package main

import "fmt"

func main() {
	var message = "stunod dna maj";

	fmt.Println(reverseWords(message));

}

func reverseWords(message string) string {
	message2 := []rune(message)
	for i, j := 0, len(message2) - 1; i < j; i, j = i+1, j-1 {
		message2[i], message2[j] = message2[j], message2[i]
	}
	fmt.Println(len(message2))

	return string(message2)
}
