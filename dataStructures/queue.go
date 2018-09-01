package main

import (
    "fmt"
)

func main() {
    queue := make([]int, 0)
    // Push to the queue
    queue = append(queue, 1)
    // Top (just get next element, don't remove it)
    fmt.Println(queue[0])
    // Discard top element
    queue = queue[1:]
    // Is empty ?
    if len(queue) == 0 {
        fmt.Println("Queue is empty !")
    }
}
