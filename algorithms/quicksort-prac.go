package main

import (
	"fmt"
)

func main() {
	var s = []int{12,4,3,6,6,2,9,2}
	fmt.Println(quicksort(s))
}
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := len(a)/2

	a[right], a[pivot] = a[pivot], a[right]
	
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])
	
	return a
}