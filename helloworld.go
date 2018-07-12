 package main

import (
	"fmt"
)
//general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming
func main() {
	fmt.Println("Hellosd World!")
	// initializes to zero
	var x int
	fmt.Println(x)

	//declaring multiple variables
	var width, height int = 100, 50 
	fmt.Println("width is", width, "height is", height)
	
	//declare variables belonging to different types in a single statement
	var (
		age = 25
		height2 = "161cm"
	)
	fmt.Println(age, height2)

	//short hand declaration and it uses := operator.
	//Short hand syntax can only be used when at least one of the variables in the left side of := is newly declared
	km := 0
	fmt.Println("  ")
	fmt.Println(km)
	fmt.Println("  ")

	// 	The following are the basic types available in go

	// bool
	// Numeric Types
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	// float32, float64
	// complex64, complex128
	// byte
	// rune
	// string

	//boolean    
	a := true
    b := false
    fmt.Println("a:", a, "b:", b)
    c := a && b
    fmt.Println("c:", c)
    d := a || b
    fmt.Println("d:", d)

}
