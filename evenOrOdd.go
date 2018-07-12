package main

 import (
         "fmt"
 )

 func main(){
	i := 0
	for i <= 5 {
		if i % 2 == 0 {
			fmt.Println("EVEN", i)
		} else {
			fmt.Println("ODD", i)
		}
		i++
	}
 }