package main

import "fmt"

func main() {
	c := make(chan bool, 2)
	c <- true
	c <- false
	c <- false
	fmt.Println(<-c)
	fmt.Println(<-c)

}
