//Implement a method to perform basic string compression using the counts of repeated characters. For example, the string a a b c c c c c a a a would become a2blc5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string.

package chapterone

import "fmt"

func compressRepeatedCharacters(input string) string {

	chars := []rune(input)
	compressed := make(map[rune]int)

	for _, v := range chars {
		compressed[v]++
	}
	var str string
	for key, val := range compressed {

		str = fmt.Sprintf("%v%d", str+string(key), val)
	}
	return str
}
