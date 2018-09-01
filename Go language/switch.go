package main

import (
    "fmt"
    "runtime"
)

func main() {
    switch os := runtime.GOOS; os{
        case "mac":
            fmt.Println("MAC?")
        default:
            fmt.Println(os)
    }
}