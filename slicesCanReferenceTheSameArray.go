package main

import (
	"fmt"
	// "math/rand"
)

func negate(s []int) {
	for i := range s {
		s[i] = -s[i]
	}
}

func main(){
	//Slicing a slice:
    var a = []int{1,2,3,4}
	b := a[1:]
	b[0] = 0
	negate(b)
	fmt.Println(a,b)
	
	c := [10]int{}
	for _,c := range c{
		fmt.Println(c)
	}
}