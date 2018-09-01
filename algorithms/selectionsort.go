/*
n2. ( two loops)
left to right, find minimum.
heapsort better than selection sort
Heapsort greatly improves the basic algorithm by using an implicit 
heap data structure to speed up finding and removing the lowest datum.
If implemented correctly, the heap will allow finding the next lowest
element in Θ(log n) time instead of Θ(n) for the inner loop in normal 
selection sort, reducing the total running time to Θ(n log n).
*
*/
package main

import (
	"fmt"
)

//We want to set min to the index of the minimum. Not the value.
func main(){
	var arr = []int{6,5,4,3,2,1}

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				min = j
			}
			
			if(arr[i] != arr[min]){
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
	}
	fmt.Println(arr)
	
}




