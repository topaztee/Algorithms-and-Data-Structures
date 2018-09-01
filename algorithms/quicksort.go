//https://www.youtube.com/watch?v=ZHVk2blR45Q&index=83&list=PLpPXw4zFa0uKKhaSz87IowJnOTzh9tiBk

// Quick Sort in Golang
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
 
    // slice := generateSlice(5)
    slice := []int{10,7,12,6,3,2,8}
	
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    quicksort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
 
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(20) 
    }
    return slice
}

//1. Move random pivot to the end of the list.
// Have two counters. One remembers the last swap positions
//and one counter that we increment through the list
// (in this case, i is the index and left is the index with last swap position)
//If its a random list,you can  make your pivot the last element, since it just as
//likely to be the middle number in the list.

func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
      
    left, right := 0, len(a)-1
      
    // pivot := rand.Int() % len(a)
    pivot := len(a)/2
    fmt.Println("pivot..", a[pivot])
      
    a[pivot], a[right] = a[right], a[pivot]
    fmt.Println(a)
      
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            // fmt.Println("..",a)
            left++
        }
    }
     //The last thing we do is take our pivot and swap it with to where the counter is
    a[left], a[right] = a[right], a[left]
      
    quicksort(a[:left])
    quicksort(a[left+1:])
      
    return a
}
