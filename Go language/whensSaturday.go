package main

import (
    "fmt"
	"time"
	"math/rand"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(9)
	fmt.Println(randNum)
    today:= time.Now().Weekday()
    switch time.Saturday {
        case today + 0:
            fmt.Println("Today")
        case today + 1:
            fmt.Println("Tomorrow")
        case today + 2:
            fmt.Println("The day after tomorrow")
        default:
            fmt.Println("too far away")
    }
    
}