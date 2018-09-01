/*
*
*/

package main

import (
	"fmt"
)

func main() {
	//
	// unsorted := []int{5,4,3,2,1}
	d := []int{1,2,3,4,5,6}

	binarySearch(d,3)
}

func binarySearch(sortedList []int, lookingFor int) int {
	var lo int = 0
	var hi int = len(sortedList) - 1
  
	for lo <= hi {
	  var mid int = lo + (hi-lo)/2
	  var midValue int = sortedList[mid]
		
	  if midValue == lookingFor {
			fmt.Println("Middle index is:", mid)
		return mid
	  } else if midValue > lookingFor {
		// We want to use the left half of our list
		hi = mid - 1
	  } else {
		// We want to use the right half of our list
		lo = mid + 1
	  }
	}
  
	// If we get here we tried to look at an invalid sub-list
	// which means the number isn't in our list.
	return -1
  }
