
type DictionaryWords struct {
	wordInAlphaOrder, word string
}{"abd", "bad"}

func getAnagrams(s string) {
	original := "dab"
	modified := []rune(original)
	strs := []string(modified)
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
}
