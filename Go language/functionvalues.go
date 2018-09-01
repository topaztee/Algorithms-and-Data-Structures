package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3,4)
}
func main(){
    fmt.Println(compute(math.Pow)

}