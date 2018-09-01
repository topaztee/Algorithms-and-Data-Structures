package main

import (
	"fmt"
)

func main(){
	var i = []int{3,4,5,6,7,1,2}

	bs(i, 1)
}

func bs(i []int, lookingfor int) int {
	// var hi int = len(i) - 1
	// var lo int = 0

	// for lo <= hi {
	// 	mid := lo + (hi-lo)/2
	// 	midValue := i[mid]
	// 	fmt.Println("Middle index is:", midValue)
	// 	if midValue == lookingfor{
	// 		fmt.Println("Middle index is:", mid)
	// 		return mid
	// 	}else if midValue > lookingfor {
	// 		hi = mid - 1
	// 	} else {
	// 		lo = mid + 1
	// 	}
	// }
	// return -1

		mid := lo + (hi-lo)/2
}

//
abc
cab, acb, abc, bac,acb, bac, acb