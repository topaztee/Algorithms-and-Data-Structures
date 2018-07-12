package main

import (
	"fmt"
)

func main() {
	//
	// unsorted := []int{5,4,3,2,1}
	d := []int{1,2,3,4,5,6}
	d = append(d,7)
	// fmt.Println(unsorted)

	// fmt.Println(len(unsorted))

	// for i := 0; i < 	
	fmt.Println("?")
	search(d,6)
}

func search(sortedArray []int, el int) int {
	init, end := 0, len(sortedArray)-1

	for init <= end {
		middle := ((end - init) >> 1) + init

		if sortedArray[middle] == el {
fmt.Println(middle)
			return middle
		}

		if sortedArray[middle] < el {
			fmt.Println(init)
			init = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}
