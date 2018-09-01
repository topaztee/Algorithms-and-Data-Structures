package main

import "fmt"

type Vertex struct {
	x,y int
}
var x map[string]Vertex

func main(){
	x = make(map[string]Vertex)

	x["Bells"] = Vertex{
		44,55,
	}

	fmt.Println(x["Bells"])

	v, i := x["Bells"]
	fmt.Println("The value:", v, "Present?", i)
	
}