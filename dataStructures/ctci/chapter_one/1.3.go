//Given two strings, write a method to decide if one is a permutation of the other.
//definition of an anagram—two words with the same character count
package chapterone

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

//Questions, Is the comparison case sentitive and is whitespace significant

// method 1: compare the two strings
// Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
func permutationMethod1(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = SortString(str1)
	str2 = SortString(str2)
	fmt.Println(str1, str2)

	return strings.EqualFold(str1, str2)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//method 2  Solution #2: Check if the two strings have identical character count
//If you look at the source code for reflect.DeepEqual's Map case, you'll see that it first checks if both maps are nil, then it checks if they have the same length before finally checking to see if they have the same set of (key, value) pairs.
//Because reflect.DeepEqual takes an interface type, it will work on any valid map (map[string]bool, map[struct{}]interface{}, etc). Note that it will also work on non-map values, so be careful that what you're passing to it are really two maps.
//If you pass it two integers, it will happily tell you whether they are equal.
func permutationMethod2(str1, str2 string) bool {
	m1 := map[rune]int{}
	m2 := make(map[rune]int)
	wordsArr, wordsArr2 := []rune(str2), []rune(str1)

	for _, v := range wordsArr {
		m1[v]++
	}
	for _, v := range wordsArr2 {
		m2[v]++
	}

	return reflect.DeepEqual(m1, m2)

}
