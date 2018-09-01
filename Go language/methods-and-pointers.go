package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v *Vertex) ScaleMethod(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}	

func ScaleFunc(v *Vertex, f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	(&v).ScaleMethod(2) // == v.ScaleMethod(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.ScaleMethod(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
