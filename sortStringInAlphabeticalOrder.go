package main

import (
	"fmt"
	"sort"
)

func main() {
    w1 := "bcad"
    fmt.Println(SortStringByCharacter(w1))
}


func SortStringByCharacter(s string) string {

	r := []rune(s)
	
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
    })

	return string(r)	
}