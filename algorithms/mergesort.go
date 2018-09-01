package main

import (
	"fmt"
)

func Merge(a, b []int) []int {
	ret := make([]int, 0, len(a)+len(b))
	for len(a) > 0 || len(b) > 0 {
		if len(a) == 0 {
			return append(ret, b...)
		}
		if len(b) == 0 {
			return append(ret, a...)
		}
		if a[0] <= b[0] {
			ret = append(ret, a[0])
			a = a[1:]
			} else {
				ret = append(ret, b[0])
				b = b[1:]
			}
		}
		
	return ret
}
//divide
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	half := len(s) / 2
	l := MergeSort(s[:half])
	r := MergeSort(s[half:])
	return Merge(l, r)
}

func main() {
	s := []int{8,7,6,5,4,3,2,1}
	fmt.Printf("%v\n%v\n....", s, MergeSort(s))
	// MergeSort(s)
}