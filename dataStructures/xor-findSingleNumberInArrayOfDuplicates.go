package main

import (
	"fmt"
)

func main() {
	array := []int{9,9,8,7,7}
	val := 0
	for i := 0; i < len(array); i++ {        
		val ^= array[i];   
	}    
	fmt.Println(val)
	
}
